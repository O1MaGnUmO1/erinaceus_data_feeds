package logpoller

import (
	"context"
	"erinaceus_data_feeds/client"
	aggregator "erinaceus_data_feeds/contract"
	diffchecker "erinaceus_data_feeds/diffChecker"
	"erinaceus_data_feeds/services/timer"
	wallet_service "erinaceus_data_feeds/services/wallet"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

// CustomFormatter is a Logrus formatter that adds two newline characters to log entries.
type CustomFormatter struct {
	logrus.JSONFormatter
}

// Format formats the log entry and adds two newline characters to the end.
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data, err := f.JSONFormatter.Format(entry)
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

type LogPoller struct {
	client          *client.Client
	contractAddress common.Address
	replayFromBlock uint64
	fromPoller      bool
	pollTicker      time.Ticker
	eventSignatures []common.Hash
	ReplayFromBlock uint64
	NewHeadCh       chan uint64
	walletService   *wallet_service.WalletService
	logchanel       chan *aggregator.AggregatorNewRound
	pendingRound    uint32
	latestAnswer    float64
	Mu              sync.Mutex
	logger          *logrus.Logger
	aggregator      *aggregator.Aggregator
	timer           *timer.Timer
}

func NewLogPoller(client *client.Client, replayFromBlock uint64, contractAddress common.Address, walletService *wallet_service.WalletService, timer *timer.Timer) (*LogPoller, error) {
	parsedABI, err := abi.JSON(strings.NewReader(aggregator.AggregatorABI))
	if err != nil {
		return nil, err
	}

	aggregatorContract, err := aggregator.NewAggregator(contractAddress, client.EthClient)
	if err != nil {
		return &LogPoller{}, fmt.Errorf("failed to create aggregator instance %v", err)
	}

	newRoundSig := parsedABI.Events["NewRound"].ID
	if newRoundSig == (common.Hash{}) {
		return nil, fmt.Errorf("event 'NewRound' not found in contract ABI")
	}

	answerUpdatedSig := parsedABI.Events["AnswerUpdated"].ID
	if answerUpdatedSig == (common.Hash{}) {
		return nil, fmt.Errorf("event 'AnswerUpdated' not found in contract ABI")
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})

	logger.Info("Starting log poller ...")
	time.Sleep(1 * time.Second)

	return &LogPoller{
		client:          client,
		contractAddress: contractAddress,

		replayFromBlock: replayFromBlock,
		logger:          logger,
		fromPoller:      false,
		pendingRound:    uint32(0),
		walletService:   walletService,
		NewHeadCh:       make(chan uint64),
		latestAnswer:    0.0,
		pollTicker:      *time.NewTicker(30 * time.Second),
		logchanel:       make(chan *aggregator.AggregatorNewRound),
		aggregator:      aggregatorContract,
		eventSignatures: []common.Hash{newRoundSig, answerUpdatedSig},
		timer:           timer,
	}, nil
}

func (lp *LogPoller) PollLogs() error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{lp.contractAddress},
		Topics:    [][]common.Hash{lp.eventSignatures},
		FromBlock: new(big.Int).SetUint64(lp.ReplayFromBlock),
		ToBlock:   nil,
	}
	lb, err := lp.client.EthClient.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get latest block %v", err)
	}
	lp.replayFromBlock = lb
	logs, err := lp.client.FilterLogs(context.Background(), query)
	if err != nil {
		errorMsg := fmt.Sprintf("error filter logs reason : %s", err)
		lp.logger.Errorf(errorMsg)
		return fmt.Errorf(errorMsg)
	}
	recentRoundId, err := lp.aggregator.LatestRound(nil)
	if err != nil {
		return fmt.Errorf("failed to get latest round Id %v", err)
	}

	for _, log := range logs {
		newRound, err := lp.aggregator.ParseNewRound(log)
		if err != nil {
			continue
		}
		if newRound.RoundId.Cmp(recentRoundId) != -1 {
			lp.logchanel <- newRound
		}
	}
	return nil
}

func (lp *LogPoller) StartPollingLogs() {
	for {
		<-lp.pollTicker.C
		lp.PollLogs()
	}
}

func (lp *LogPoller) StartListeningForPrices() {
	for {
		select {
		case newRound := <-lp.logchanel:
			lp.logger.WithFields(logrus.Fields{
				"Started By": newRound.StartedBy,
				"Started At": newRound.StartedAt,
				"RoundID":    newRound.RoundId,
			}).Info("Received new round request, trying to submit ...")
			if newRound.StartedBy == lp.walletService.Key.Address {
				lp.logger.Info("log is our own, skiping ...")
				continue
			}
			currentAnswer, err := lp.aggregator.LatestRoundData(nil)
			if err != nil {
				lp.logger.Errorf("failed to get latest round data %v", err)
				continue
			}
			nextAnswer, err := lp.timer.FetchData()
			if err != nil {
				lp.logger.Errorf("failed to make http request %v", err)
				continue
			}
			next := new(big.Int).SetUint64(uint64(nextAnswer * 100))
			if diffchecker.CheckDifference(currentAnswer.Answer, next) {
				lp.logger.WithFields(logrus.Fields{
					"Current Answer": currentAnswer,
					"Next Answer":    next,
				}).Info("Met difference Submitting ...")
				if err := lp.TrySubmit(0, next); err != nil {
					lp.logger.Errorf("failed to submit difference %v", err)
					continue
				}
				continue
			}
			lp.logger.WithFields(logrus.Fields{
				"Round started by": newRound.StartedBy,
				"Our Address":      lp.walletService.Key.Address,
			}).Info()


			if err := lp.TrySubmit(uint32(newRound.RoundId.Uint64()), next); err != nil {
				lp.logger.Errorf("failed to answer %v", err)
				continue
			}
		case head := <-lp.NewHeadCh:
			query := ethereum.FilterQuery{
				Addresses: []common.Address{lp.contractAddress},
				Topics:    [][]common.Hash{lp.eventSignatures},
				FromBlock: new(big.Int).SetUint64(lp.ReplayFromBlock),
				ToBlock:   new(big.Int).SetUint64(head),
			}

		case price := <-lp.timer.DrumbeatChan:
			lp.logger.WithFields(logrus.Fields{
				"Answer":    price,
				"Timestamp": time.Now().UTC(),
			}).Info("Received price with 2m interval")
			nextAnswer, err := lp.timer.FetchData()
			if err != nil {
				lp.logger.Errorf("failed to make http request %v", err)
				continue
			}
			next := new(big.Int).SetUint64(uint64(nextAnswer * 100))
			if err := lp.TrySubmit(0, next); err != nil {
				lp.logger.Errorf("failed to answer after 2 min %v", err)
				continue
			}
		}
	}
}

func (lp *LogPoller) checkOurNewRoundLog() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{lp.contractAddress},
		Topics:    [][]common.Hash{lp.eventSignatures},
		FromBlock: new(big.Int).SetUint64(lp.ReplayFromBlock),
		ToBlock: nil,
	}
	logs, err := lp.client.FilterLogs(context.Background(), query)
	if err != nil {
		errorMsg := fmt.Sprintf("error filter logs reason : %s", err)
		lp.logger.Errorf(errorMsg)
		return
	}
	for _, log := range logs {
		newRound, err := lp.aggregator.ParseNewRound(log)
		if err != nil {
			continue
		}
		if newRound.StartedBy == lp.walletService.Key.Address {
			lp.
		}
	}
}

func (lp *LogPoller) TrySubmit(roundId uint32, answer *big.Int) error {
	timeNow := time.Now()
	lp.timer.Ticker1.Stop()
	roundState, err := lp.aggregator.OracleRoundState(nil, lp.walletService.Key.Address, roundId)
	if err != nil {
		return fmt.Errorf("failed to get oracle sound state %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(lp.walletService.Key.ToEcdsaPrivKey(), big.NewInt(4090))
	if err != nil {
		return fmt.Errorf("failed to create keyed transactor %v", err)
	}
	if roundState.EligibleToSubmit {
		tx, err := lp.aggregator.Submit(auth, new(big.Int).SetUint64(uint64(roundState.RoundId)), answer)
		if err != nil {
			return fmt.Errorf("failed to submit %v", err)
		}
		lp.logger.WithFields(logrus.Fields{
			"Tx":        tx,
			"Timestamp": time.Now().UTC(),
		}).Info("Trying to send transaction")

		receipt, err := bind.WaitMined(context.Background(), lp.client.EthClient, tx)
		if err != nil {
			return fmt.Errorf("failed to wait tx to be mined %v", err)
		}

		if receipt.Status == 0x1 {
			lp.logger.WithFields(logrus.Fields{
				"Receipt":   receipt,
				"Timestamp": time.Now().UTC(),
			}).Info("Transaction successfully sent")
			lp.pollTicker.Reset(30 * time.Second)
			lp.timer.Ticker1.Reset(2 * time.Minute)
		}
	} else {
		dur := time.Since(timeNow) + 2*time.Minute
		lp.timer.Ticker1.Reset(dur)
		return fmt.Errorf("not eligible to submit tx")
	}
	return nil
}

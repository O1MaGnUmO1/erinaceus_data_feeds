package logpoller

import (
	"context"
	"erinaceus_data_feeds/client"
	aggregator "erinaceus_data_feeds/contract"
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
	client            *client.Client
	contractAddress   common.Address
	replayFromBlock   uint64
	eventSignatures   []common.Hash
	ReplayFromBlock   uint64
	pendingRound      uint64
	latestBlockNumber uint64
	Mu                sync.Mutex
	logger            *logrus.Logger
	aggregator        *aggregator.Aggregator
	timer             *timer.Timer
}

func NewLogPoller(client *client.Client, replayFromBlock uint64, contractAddress common.Address, walletService *wallet_service.WalletService, timer *timer.Timer) (*LogPoller, error) {
	parsedABI, err := abi.JSON(strings.NewReader(aggregator.AggregatorABI))
	if err != nil {
		return nil, err
	}

	aggregator, err := aggregator.NewAggregator(contractAddress, client.EthClient)
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
		aggregator:      aggregator,
		eventSignatures: []common.Hash{newRoundSig, answerUpdatedSig},
		timer:           timer,
	}, nil
}

func (lp *LogPoller) PollLogs() error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{lp.contractAddress},
		Topics:    [][]common.Hash{lp.eventSignatures},
		FromBlock: new(big.Int).SetUint64(lp.ReplayFromBlock),
		ToBlock:   new(big.Int).SetUint64(lp.GetLatestBlockNumber()),
	}

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
	lp.logger.Infof("New Round Id %d", recentRoundId)
	lp.logger.Infof("got %d request from %d to %d blocks", len(logs), lp.replayFromBlock, lp.GetLatestBlockNumber())

	for _, log := range logs {
		newRound, err := lp.aggregator.ParseNewRound(log)
		if err != nil {
			continue
		}
		lp.logger.Infof("Round id inLog is %d", newRound.RoundId)
		if newRound.RoundId.Cmp(recentRoundId) != -1 {
			lp.logger.Info("got new round request")
		}
	}
	// lp.logger.Infof("Got %d NewRound logs", len(arr))
	return nil
}

func (lp *LogPoller) StartListeningForPrices() {
	for {
		select {
		case price := <-lp.timer.PriceChan:
			// Use the price
			lp.logger.Infof("Received price: %f", price)
			lp.timer.Ticker.Reset(10 * time.Second)
			// Add more cases as needed, for example, a case to handle termination

		}
	}
}

func (lp *LogPoller) SetLatestBlockNumber(blockNumber uint64) {
	lp.Mu.Lock()
	defer lp.Mu.Unlock()
	lp.latestBlockNumber = blockNumber
}

// getLastProcessedBlock returns the last processed block number
func (lp *LogPoller) GetLatestBlockNumber() uint64 {
	return lp.latestBlockNumber
}

// func (lp *LogPoller) getFromCache(key string) (*erinaceus_vrf.ErinacuesVrfRandomWordsRequested, bool) {
// 	lp.Mu.Lock()
// 	defer lp.Mu.Unlock()
// 	value, ok := lp.Unfulfilled[key]
// 	return value, ok
// }

// func (lp *LogPoller) addToCache(key string, value *erinaceus_vrf.ErinacuesVrfRandomWordsRequested) {
// 	lp.Mu.Lock()
// 	defer lp.Mu.Unlock()
// 	lp.Unfulfilled[key] = value
// }

// func (lp *LogPoller) deleteFromCache(key string) {
// 	lp.Mu.Lock()
// 	defer lp.Mu.Unlock()
// 	_, ok := lp.Unfulfilled[key]
// 	if ok {
// 		delete(lp.Unfulfilled, key)
// 	}
// }

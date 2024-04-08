package application

import (
	"context"
	"erinaceus_data_feeds/client"
	"erinaceus_data_feeds/headtracker"
	logpoller "erinaceus_data_feeds/logPoller"
	"erinaceus_data_feeds/services/timer"
	wallet_service "erinaceus_data_feeds/services/wallet"
	"erinaceus_data_feeds/utils"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var replayFromBlock = uint64(1000000) // default number of replay blocks
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

type Application struct {
	Client        *client.Client
	LogPoller     *logpoller.LogPoller
	WalletService *wallet_service.WalletService
	HeadTracker   *headtracker.HeadTracker
	Logger        *logrus.Logger
}

func NewApplication() (*Application, error) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})
	wd, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("failed to get working directory %v", err)
		return nil, err
	}
	useConfig := flag.Bool("with-config", false, "Set to true to use configuration file")
	flag.Parse()
	if *useConfig {
		if err := utils.LoadEnv(filepath.Join(wd, "/.env")); err != nil {
			return nil, fmt.Errorf("failed to set env variables from .env %v", err)
		}
	}
	client, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create application : Err=<%v>", err)
	}
	// replayFromBlock := replayFromBlock
	// if os.Getenv("EC_REPLAY_FROM_BLOCK") != "" {
	// 	replayFromBlock, err = strconv.ParseUint(os.Getenv("EC_REPLAY_FROM_BLOCK"), 10, 64)
	// 	if err != nil {
	// 		logrus.Errorf("failed to get replay block number will use deafult value %v", err)
	// 	}
	// }
	walletService := wallet_service.NewWalletService(client)
	contractAddress := common.HexToAddress("0x318E735e110265044508C5DC95B38C1074D43829")
	timer, _ := timer.NewTimerService()

	logpoller, err := logpoller.NewLogPoller(client, replayFromBlock, contractAddress, walletService, timer)

	if err != nil {
		return nil, fmt.Errorf("failed to create log poller : Err=<%v>", err)
	}
	go timer.Start()
	go logpoller.StartListeningForPrices()

	headtracker := headtracker.NewHeadTracker(client, logpoller)
	return &Application{
		Client:        client,
		LogPoller:     logpoller,
		WalletService: walletService,
		HeadTracker:   headtracker,
		Logger:        logger,
	}, nil
}

func (app *Application) Run() {
	_, err := app.WalletService.CreateNewFTNKey()
	if err != nil {
		app.Logger.Errorf("failed to create FTN Key %v", err)
	}
	app.WalletService.PrintWalletDetails()

	err = app.HeadTracker.Start(context.Background())
	if err != nil {
		fmt.Printf("error starting head tracker %v", err)
	}
}

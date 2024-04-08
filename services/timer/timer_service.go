package timer

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var interval = time.Second * 24

type CustomFormatter struct {
	logrus.JSONFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data, err := f.JSONFormatter.Format(entry)
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

type PriceResponse struct {
	Ripple struct {
		Usd float64 `json:"usd"`
	} `json:"ripple"`
}

type Timer struct {
	interval  time.Duration
	url       string
	Ticker    *time.Ticker
	logger    *logrus.Logger
	PriceChan chan float64
}

func NewTimerService() (*Timer, error) {
	// interval := os.Getenv("EC_UPDATE_INTERVAL")
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})
	return &Timer{
		interval:  interval,
		url:       "http://localhost:8080/api/v3/simple/price?ids=ripple&vs_currencies=usd",
		Ticker:    time.NewTicker(interval),
		logger:    logger,
		PriceChan: make(chan float64),
	}, nil
}

func (t *Timer) Start() {
	t.logger.WithFields(logrus.Fields{
		"Timestamp":      time.Now().UTC(),
		"Timer Interval": t.interval.String(),
	}).Info("Starting Timer Service")
	for {
		<-t.Ticker.C // Wait for the next tick
		t.logger.WithFields(logrus.Fields{
			"URL": t.url,
		}).Info("Making Http Get Request")
		response, err := http.Get(t.url)
		if err != nil {
			logrus.Errorf("error making request: %v", err)
			continue
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorf("error reading response: %v", err)
			continue
		}

		var resp PriceResponse
		// logrus.Infof("Received response: %s\n", body)
		err = json.Unmarshal(body, &resp)
		if err != nil {
			logrus.Errorf("failed to unmarshall %v, Err: %v", response, err)
		}

		t.logger.WithFields(logrus.Fields{
			"Timestamp": time.Now().UTC(),
			"Answer":    resp.Ripple.Usd,
		}).Infof("Got Response from %s", t.url)
		t.PriceChan <- resp.Ripple.Usd

		response.Body.Close()
	}
}

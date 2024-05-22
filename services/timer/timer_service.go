package timer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

var interval = time.Second * 30

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

// APIRequestDetails encapsulates details for making API requests
type APIRequestDetails struct {
	URL      string
	Headers  map[string]string
	JSONPath string // JSON path to extract the price
}

func NewAPIRequestDetails() *APIRequestDetails {
	api := os.Getenv("EC_API_URL")
	header := os.Getenv("EC_API_HEADER_NAME")
	apiKey := os.Getenv("EC_API_KEY")

	headersMap := make(map[string]string)
	if header != "" && apiKey != "" {
		headersMap[header] = apiKey
	}

	jsonPath := os.Getenv("EC_API_JSON_PATH")

	return &APIRequestDetails{
		URL:      api,
		Headers:  headersMap,
		JSONPath: jsonPath,
	}
}

type Timer struct {
	interval     time.Duration
	apiDetails   *APIRequestDetails
	Ticker       *time.Ticker
	Ticker1      *time.Ticker
	logger       *logrus.Logger
	PriceChan    chan float64
	DrumbeatChan chan float64
}

func NewTimerService() (*Timer, error) {
	// interval := os.Getenv("EC_UPDATE_INTERVAL")
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&CustomFormatter{})
	return &Timer{
		interval:     interval,
		Ticker1:      time.NewTicker(2 * time.Minute),
		logger:       logger,
		apiDetails:   NewAPIRequestDetails(),
		PriceChan:    make(chan float64),
		DrumbeatChan: make(chan float64),
	}, nil
}

func (t *Timer) Start() {
	t.logger.WithFields(logrus.Fields{
		"Timestamp":      time.Now().UTC(),
		"Timer Interval": t.interval.String(),
	}).Info("Starting Timer Service")
	for {
		select {
		case <-t.Ticker1.C:
			resp, err := t.FetchData()
			if err != nil {
				logrus.Errorf("failed to make http request %v", err)
				continue
			}
			t.DrumbeatChan <- resp
		}
	}
}

func (t *Timer) FetchData() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", t.apiDetails.URL, nil)
	if err != nil {
		return 0.0, err
	}

	// Add headers to the request
	for key, value := range t.apiDetails.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	// Use gjson to parse and extract the value from the JSON dynamically
	result := gjson.GetBytes(body, t.apiDetails.JSONPath)
	if !result.Exists() {
		return 0.0, fmt.Errorf("failed to extract price using JSONPath: %s", t.apiDetails.JSONPath)
	}

	return result.Float(), nil
}

// func (t *Timer) MakeHttpRequest() (float64, error) {
// 	t.logger.WithFields(logrus.Fields{
// 		"URL": t.url,
// 	}).Info("Making Http Get Request")
// 	response, err := http.Get(t.url)
// 	if err != nil {
// 		return 0.0, fmt.Errorf("error making request: %v", err)
// 	}
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return 0.0, fmt.Errorf("error reading response: %v", err)
// 	}

// 	var resp PriceResponse
// 	err = json.Unmarshal(body, &resp)
// 	if err != nil {
// 		logrus.Errorf("failed to unmarshall %v, Err: %v", response, err)
// 	}

// 	response.Body.Close()
// 	return resp.Ripple.Usd, nil
// }

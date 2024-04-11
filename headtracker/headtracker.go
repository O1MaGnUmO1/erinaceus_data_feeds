package headtracker

// import (
// 	"context"
// 	"erinaceus_data_feeds/client"
// 	logpoller "erinaceus_data_feeds/logPoller"
// 	"os"
// 	"sync"
// 	"time"

// 	"github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/sirupsen/logrus"
// )

// var retryInterval = 5 * time.Second

// type CustomFormatter struct {
// 	logrus.JSONFormatter
// }

// func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
// 	data, err := f.JSONFormatter.Format(entry)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return append(data, '\n'), nil
// }

// type HeadTracker struct {
// 	client    *client.Client
// 	headers   chan *types.Header
// 	logger    *logrus.Logger
// 	sub       ethereum.Subscription
// 	logPoller *logpoller.LogPoller
// 	wg        sync.WaitGroup
// }

// func NewHeadTracker(client *client.Client, logPoller *logpoller.LogPoller) *HeadTracker {
// 	logger := logrus.New()
// 	logger.SetOutput(os.Stdout)
// 	logger.SetFormatter(&CustomFormatter{})
// 	return &HeadTracker{
// 		client:    client,
// 		headers:   make(chan *types.Header),
// 		logger:    logger,
// 		logPoller: logPoller,
// 		wg:        sync.WaitGroup{},
// 	}
// }

// func (ht *HeadTracker) Start(ctx context.Context) error {
// 	err := ht.subscribeToNewHead(ctx)
// 	if err != nil {
// 		ht.logger.Errorf("Failed to start subscription: %v", err)
// 		return err // Return here or decide to keep retrying based on your logic.
// 	}

// 	ht.wg.Add(1)
// 	defer ht.wg.Done()
// 	go ht.handleSubscription(ctx)

// 	// Wait for the subscription handler to finish (it may never finish if context is not canceled)
// 	ht.wg.Wait()
// 	return nil
// }

// func (ht *HeadTracker) subscribeToNewHead(ctx context.Context) error {
// 	sub, err := ht.client.SubscribeNewHead(ctx, ht.headers)
// 	if err != nil {
// 		ht.logger.Errorf("Failed to subscribe to new head: %v", err)
// 		return err
// 	}
// 	ht.sub = sub
// 	ht.logger.Infoln("Successfully subscribed to new head")
// 	return nil
// }

// func (ht *HeadTracker) handleSubscription(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			if ht.sub != nil {
// 				ht.sub.Unsubscribe()
// 			}
// 			return
// 		case err := <-ht.sub.Err():
// 			ht.logger.Errorf("Subscription error: %v", err)
// 			time.Sleep(retryInterval) // Wait before retrying to avoid spamming the node

// 			for {
// 				// Attempt to resubscribe until successful or context is cancelled
// 				if err := ht.subscribeToNewHead(ctx); err != nil {
// 					ht.logger.Errorf("Failed to resubscribe: %v", err)
// 					time.Sleep(retryInterval) // Wait before retrying
// 					select {
// 					case <-ctx.Done(): // Check if context was cancelled during retry wait
// 						return
// 					default:
// 						continue // Retry subscription
// 					}
// 				} else {
// 					break // Successfully resubscribed
// 				}
// 			}
// 		case header := <-ht.headers:
// 			ht.processNewHead(header)
// 		}
// 	}
// }

// func (ht *HeadTracker) processNewHead(header *types.Header) {
// 	ht.logger.WithFields(logrus.Fields{
// 		"Head Number": header.Number,
// 		"Timestamp":   header.Time,
// 		"Hash":        header.Hash(),
// 		"Parent Hash": header.ParentHash,
// 		"Gas Limit":   header.GasLimit,
// 		"Gas Used":    header.GasUsed,
// 	}).Infoln("Received new head")
// 	ht.logPoller.SetLatestBlockNumber(header.Number.Uint64())

// 	if err := ht.logPoller.PollLogs(); err != nil {
// 		ht.logger.Errorf("Error processing requests: %v", err)
// 	}
// }

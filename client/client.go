package client

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	EthClient *ethclient.Client
}

func (cl *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return cl.EthClient.FilterLogs(ctx, query)
}

func (cl *Client) BalanceAt(ctx context.Context, address common.Address, blockNum *big.Int) (*big.Int, error) {
	return cl.EthClient.BalanceAt(ctx, address, blockNum)
}

func (cl *Client) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return cl.EthClient.PendingNonceAt(ctx, address)
}

func (cl *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return cl.EthClient.SendTransaction(ctx, tx)
}

func (cl *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return cl.EthClient.EstimateGas(ctx, msg)
}

func (cl *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return cl.EthClient.SubscribeNewHead(ctx, ch)
}

func NewClient() (*Client, error) {
	client, err := ethclient.Dial(os.Getenv("EC_NODE_URL"))
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)

		return nil, err
	}
	return &Client{
		EthClient: client,
	}, nil
}

package wallet_service

import (
	"context"
	"erinaceus_data_feeds/client"
	"erinaceus_data_feeds/keys/ethkey"
	"erinaceus_data_feeds/utils"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var basepath, _ = os.Getwd()

type WalletService struct {
	Client *client.Client
	Key    ethkey.KeyV2
}

func NewWalletService(client *client.Client) *WalletService {
	logrus.Info("Starting Wallet Service ....\n")
	time.Sleep(1 * time.Second)
	return &WalletService{
		Client: client,
	}
}

func (w *WalletService) CreateNewFTNKey() (ethkey.KeyV2, error) {
	ethKey, err := ethkey.GetKeyIfEnvSet()
	if err != nil {
		return ethkey.KeyV2{}, err
	}

	if ethKey.Address.Cmp(common.Address{}) != 0 {
		w.Key = ethKey
		return w.Key, nil
	}

	logrus.Info("Generating New FTN Key...")
	key, err := ethkey.NewV2()
	if err != nil {
		logrus.Errorf("Failed to generate FTN key: %v", err)
		return ethkey.KeyV2{}, err
	}

	w.Key = key
	keyJSON, err := key.ToEncryptedJSON(os.Getenv("EC_FTN_KEY_PASSWORD"), utils.FastScryptParams)
	if err != nil {
		logrus.Errorf("Failed to encrypt key: %v", err)
		return ethkey.KeyV2{}, err
	}

	keyFilePath := filepath.Join(basepath, "ftn_key.json")
	if err := os.WriteFile(keyFilePath, keyJSON, 0644); err != nil {
		logrus.Errorf("Failed to write key file: %v", err)
		return ethkey.KeyV2{}, err
	}

	logrus.Infof("File %s is saved", keyFilePath)
	return w.Key, nil
}

// weiToETH converts Wei to Ether and returns the result as a string
func WeiToETH(wei *big.Int) string {
	ether := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1000000000000000000))
	return ether.Text('f', 18) // Format with 18 decimal places
}

func (w *WalletService) PrintWalletDetails() {
	logrus.Info("Succesfully generated FTN Wallet")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("FTN Address")
	fmt.Println(w.Key.Address.Hex())
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Balance")
	balance, err := w.Client.BalanceAt(context.Background(), w.Key.Address, nil)
	if err != nil {
		logrus.Errorf("Failed to get balance: %v", err)
	}
	fmt.Println(WeiToETH(balance), "FTN")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println()
}

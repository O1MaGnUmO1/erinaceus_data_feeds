package wallet_service

import (
	"erinaceus_data_feeds/client"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestNewWalletService(t *testing.T) {
	mockClient := &client.Client{}
	walletService := NewWalletService(mockClient)

	if *walletService.Client != *mockClient {
		t.Errorf("Client was not set correctly in NewWalletService")
	}
}

func TestCreateNewFTNKey(t *testing.T) {
	// Use a temporary directory for file operations
	tmpDir := t.TempDir()
	os.Chdir(tmpDir)

	mockClient := &client.Client{} // Initialize your client mock
	walletService := NewWalletService(mockClient)

	// Clear environment variable to simulate the key not being set
	os.Unsetenv("EC_FTN_KEY_JSON")

	key, err := walletService.CreateNewFTNKey()
	if err != nil {
		t.Errorf("Error should not have occurred during CreateNewFTNKey: %v", err)
	}

	if key.Address.Cmp(common.Address{}) == 0 {
		t.Errorf("Generated key address is empty")
	}

	// Check if the file is created
	if _, err := os.Stat("ftn_key.json"); os.IsNotExist(err) {
		t.Errorf("ftn_key.json file was not created")
	}

}

func TestWeiToETH(t *testing.T) {
	testCases := []struct {
		wei string
		eth string
	}{
		{"1000000000000000000", "1.000000000000000000"},
		{"2000000000000000000", "2.000000000000000000"},
		{"1234567890000000000", "1.234567890000000000"},
	}

	for _, tc := range testCases {
		wei := new(big.Int)
		wei.SetString(tc.wei, 10)
		eth := WeiToETH(wei)
		if eth != tc.eth {
			t.Errorf("Expected %s ETH, got %s ETH for %s Wei", tc.eth, eth, tc.wei)
		}
	}
}

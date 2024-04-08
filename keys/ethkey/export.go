package ethkey

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"erinaceus_data_feeds/utils"
)

var basepath, _ = os.Getwd()

type EncryptedEthKeyExport struct {
	KeyType string              `json:"keyType"`
	Address EIP55Address        `json:"address"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
}

func (key KeyV2) ToEncryptedJSON(password string, scryptParams utils.ScryptParams) (export []byte, err error) {
	// DEV: uuid is derived directly from the address, since it is not stored internally
	id, err := uuid.FromBytes(key.Address.Bytes()[:16])
	if err != nil {
		return nil, errors.Wrapf(err, "could not generate ethkey UUID")
	}
	dKey := &keystore.Key{
		Id:         id,
		Address:    key.Address,
		PrivateKey: key.privateKey,
	}
	return keystore.EncryptKey(dKey, password, scryptParams.N, scryptParams.P)
}

func GetKeyIfEnvSet() (KeyV2, error) {
	keyFilePath := os.Getenv("EC_FTN_KEY_JSON_PATH")
	keyPassword := os.Getenv("EC_FTN_KEY_PASSWORD")

	if keyPassword == "" {
		errorMsg := "EC_FTN_KEY_PASSWORD is not set. Please set it and re-run the program"
		logrus.Error(errorMsg)
		return KeyV2{}, fmt.Errorf(errorMsg)
	}

	if keyFilePath == "" {
		keyFilePath = filepath.Join(basepath, "ftn_key.json")
	}

	if !utils.FileExists(keyFilePath) {
		return KeyV2{}, nil
	}

	jsonData, err := os.ReadFile(keyFilePath)
	if err != nil {
		logrus.Errorf("Failed to read key file: %v", err)
		return KeyV2{}, err
	}

	key, err := keystore.DecryptKey(jsonData, keyPassword)
	if err != nil {
		logrus.Errorf("Failed to decrypt key: %v", err)
		return KeyV2{}, err
	}

	return FromPrivateKey(key.PrivateKey), nil
}

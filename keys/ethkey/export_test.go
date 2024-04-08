package ethkey

import (
	"testing"

	"erinaceus_data_feeds/keys"
)

func TestEthKeys_ExportImport(t *testing.T) {
	keys.RunKeyExportImportTestcase(t, createKey, func(keyJSON []byte, password string) (kt keys.KeyType, err error) {
		t.SkipNow()
		return kt, err
	})
}

func createKey() (keys.KeyType, error) {
	return NewV2()
}

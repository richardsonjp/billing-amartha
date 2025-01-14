package enum

import (
	"go-skeleton/pkg/utils/constant"
	"strings"
)

// WalletStatus represents the status of a wallet (enabled/disabled)
type WalletStatus int64

// Scan for converting byte to string for fetching/read
func (s *WalletStatus) Scan(value interface{}) error {
	key := value.(string)
	for i, v := range walletStatusKey {
		if v == key {
			*s = i
		}
	}
	return nil
}

func NewWalletStatus(value string) WalletStatus {
	for i, v := range walletStatusKey {
		if v == value {
			return i
		}
	}
	panic("enum not found")
}

const (
	ENABLED WalletStatus = iota + 1
	DISABLED
)

var walletStatusKey = map[WalletStatus]string{
	ENABLED:  "enabled",
	DISABLED: "disabled",
}

// String for stringify WalletStatus
func (s WalletStatus) String() string {
	return walletStatusKey[s]
}

func GetWalletStatusKeyValuePairs() []constant.KeyValue {
	var arr []constant.KeyValue
	for _, v := range walletStatusKey {
		arr = append(arr, constant.KeyValue{
			Key:   v,                // act as key
			Value: strings.Title(v), // act as label
		})
	}
	return arr
}

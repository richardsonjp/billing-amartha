package enum

import (
	"go-skeleton/pkg/utils/constant"
	"strings"
)

// TransactionType is connected with model billing which needs deposit/withdrawal type
type TransactionType int64

// Scan for converting byte to string for fetching/read
func (s *TransactionType) Scan(value interface{}) error {
	key := value.(string)
	for i, v := range transactionTypeKey {
		if v == key {
			*s = i
		}
	}
	return nil
}

func NewTransactionType(value string) TransactionType {
	for i, v := range transactionTypeKey {
		if v == value {
			return i
		}
	}
	panic("enum not found")
}

const (
	DEPOSIT TransactionType = iota + 1
	WITHDRAWAL
)

var transactionTypeKey = map[TransactionType]string{
	DEPOSIT:    "deposit",
	WITHDRAWAL: "withdrawal",
}

// String for stringify TransactionType
func (s TransactionType) String() string {
	return transactionTypeKey[s]
}

func GetTransactionTypeKeyValuePairs() []constant.KeyValue {
	var arr []constant.KeyValue
	for _, v := range transactionTypeKey {
		arr = append(arr, constant.KeyValue{
			Key:   v,                // act as key
			Value: strings.Title(v), // act as label
		})
	}
	return arr
}

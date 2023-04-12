package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/vincent-petithory/dataurl"
	"math/big"
	"strings"
	"time"
)

func IsFieldNew[K int | float64 | string](newField *K, oldField K) bool {
	return newField != nil && *newField != oldField
}

func IsIntFieldNew(newField *int, oldField int) bool {
	return newField != nil && *newField != oldField
}

func IsStringFieldNew(newField *string, oldField string) bool {
	return newField != nil && *newField != oldField
}

func IsTimeFieldNew(newField *time.Time, oldField time.Time) bool {
	return newField != nil && !(*newField).Equal(oldField)
}

func IsStringNil(field *string) bool {
	return field == nil || len(*field) == 0
}

func IsIntNil(field *int) bool {
	return field == nil
}

func IsBoolFieldNew(newField *bool, oldField bool) bool {
	return newField != nil && *newField != oldField
}

// GetBytesFromDataURL parses dataURL string and retrieves bytes
func GetBytesFromDataURL(dataURLString string) ([]byte, error) {
	dataURL, err := dataurl.DecodeString(dataURLString)
	if err != nil {
		return nil, err
	}

	return dataURL.Data, nil
}

// ParseBigFloat parse string value to big.Float
func ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func IsItemInList[K string | int](item K, list []K) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}

	return false
}

func ValPointer[K bool | string | time.Time | int](val K) *K {
	return &val
}

package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/vincent-petithory/dataurl"
	"math/big"
	"math/rand"
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

// IsItemInList checks if item is in list
func IsItemInList[K string | int](item K, list []K) bool {
	for _, li := range list {
		if item == li {
			return true
		}
	}

	return false
}

// GetItemIDInList returns the index of item in list
// returns -1 if item is not in list
func GetItemIDInList[K string | int](item K, list []K) int {
	for i, li := range list {
		if item == li {
			return i
		}
	}

	return -1
}

// ValPointer returns a pointer to the value passed in, uses generics
// this is useful for setting values in structs
// example:
//
//	type MyStruct struct {
//		MyField *string
//	}
//	myStruct := MyStruct{
//		MyField: ValPointer("myValue"),
//	}
func ValPointer[K bool | string | time.Time | int](val K) *K {
	return &val
}

func IsStringLooksLikeAddress(str string) bool {
	return len(str) > 2 && strings.ToLower(str[:2]) == "0x"
}

// GetKeysFromStructMap returns the keys of a map of structs with generic key type
func GetKeysFromStructMap[K string | int](m map[K]struct{}) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func IntToUInt64Slice(src []int) []uint64 {
	var dest []uint64
	for _, v := range src {
		dest = append(dest, uint64(v))
	}
	return dest
}

func UInt64ToIntSlice(src []uint64) []int {
	var dest []int
	for _, v := range src {
		dest = append(dest, int(v))
	}
	return dest
}

func GetRandomNumberInRange(min, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return min + r1.Intn(max-min)
}

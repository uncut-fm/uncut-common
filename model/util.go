package model

import (
	"github.com/vincent-petithory/dataurl"
)

func IsIntFieldNew(newField *int, oldField int) bool {
	return newField != nil && *newField != oldField
}

func IsStringFieldNew(newField *string, oldField string) bool {
	return newField != nil && *newField != oldField
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

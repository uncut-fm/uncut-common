package model

import (
	mngmt_model "github.com/uncut-fm/uncut-management-api-2/graph/model"
	"github.com/vincent-petithory/dataurl"
)

func IsIntFieldNew(newField *int, oldField int) bool {
	return newField != nil && *newField != oldField
}

func IsStringFieldNew(newField *string, oldField string) bool {
	return newField != nil && *newField != oldField
}

func IsFloatFieldNew(newField *float64, oldField float32) bool {
	oldFieldFloat64, _ := mngmt_model.ParseFloat32ToFloat64(oldField)
	return newField != nil && *newField != oldFieldFloat64
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

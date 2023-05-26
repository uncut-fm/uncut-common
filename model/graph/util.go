package model

import (
	"encoding/json"
	"time"
)

// NewModelFromProperties creates a new model from the given properties.
// uses generics to define the type of the model
func NewModelFromProperties[K User | Wallet | NFT | NFTOwner | NFTCollection](properties map[string]interface{}) (*K, error) {
	propertiesBytes, err := json.Marshal(properties)
	if err != nil {
		return nil, err
	}

	n := new(K)

	err = json.Unmarshal(propertiesBytes, n)

	return n, err
}

// NewPropertiesMapFromModel creates a new properties map from the given model.
// uses generics to define the type of the model
func NewPropertiesMapFromModel[K User | Wallet | NFT | NFTOwner | NFTCollection](model *K) (map[string]interface{}, error) {
	modelBytes, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	var properties map[string]interface{}

	err = json.Unmarshal(modelBytes, &properties)

	// loop over all keys and remove keys with map value types
	for key, value := range properties {
		// if key name starts with capital letter, remove it
		if key[0] >= 'A' && key[0] <= 'Z' {
			delete(properties, key)
			continue
		}

		if _, ok := value.(map[string]interface{}); ok {
			delete(properties, key)
			continue
		}

		// if the value is nil interface type, remove it
		if value == nil {
			delete(properties, key)
			continue
		}
	}

	return properties, err
}

// UnixTimeToTime converts unix time to time.Time
func UnixTimeToTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

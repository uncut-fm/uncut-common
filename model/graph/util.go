package model

import "encoding/json"

// NewModelFromProperties creates a new model from the given properties.
// uses generics to define the type of the model
func NewModelFromProperties[K User | NFT | NFTOwner](properties map[string]interface{}) (*K, error) {
	propertiesBytes, err := json.Marshal(properties)
	if err != nil {
		return nil, err
	}

	n := new(K)

	err = json.Unmarshal(propertiesBytes, n)

	return n, err
}

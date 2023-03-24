package alchemy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTokenID(t *testing.T) {
	type testCase struct {
		name            string
		tokenIDInHex    string
		expectedTokenID uint
	}

	cases := []testCase{
		{
			name:            "big num",
			tokenIDInHex:    "0x0f1ddb5a3c504a40e653379cf1ca8d48de13f0c5000000000000010000000001",
			expectedTokenID: 6837445208824483060242003533008093947305993890357970191003308304382715494401,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nft := Nft{ID: struct {
				TokenID       string `json:"tokenId"`
				TokenMetadata struct {
					TokenType TokenType `json:"tokenType"`
				} `json:"tokenMetadata"`
			}(struct {
				TokenID       string
				TokenMetadata struct {
					TokenType TokenType `json:"tokenType"`
				}
			}{TokenID: c.tokenIDInHex})}

			tokenID := nft.GetTokenID()
			assert.Equal(t, c.expectedTokenID, tokenID)
		})
	}
}

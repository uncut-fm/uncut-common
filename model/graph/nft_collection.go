package model

import (
	"github.com/mindstand/gogm/v2"
	"time"
)

type NFTCollection struct {
	gogm.BaseNode

	ID              int       `gogm:"name=id,pk"`
	Name            string    `gogm:"name=name"`
	ContractAddress string    `gogm:"name=contract_address"`
	CreatorAddress  string    `gogm:"name=creator_address"`
	CreatedAt       time.Time `gogm:"name=created_at"`
	UpdatedAt       time.Time `gogm:"name=updated_at"`
	UpdatedOnBlock  int       `gogm:"name=updated_on_block"`
	Network         string    `gogm:"name=network"`
	TokenType       string    `gogm:"name=token_type"`
	Origin          string    `gogm:"name=origin"`
	NFTs            []*NFT    `gogm:"direction=outgoing;relationship=BELONGS_TO"`
}

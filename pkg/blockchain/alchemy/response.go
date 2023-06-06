package alchemy

type getTokenBalancesResponse struct {
	Result struct {
		Address       string `json:"address"`
		TokenBalances []struct {
			ContractAddress string      `json:"contractAddress"`
			TokenBalance    string      `json:"tokenBalance"`
			Error           interface{} `json:"error"`
		} `json:"tokenBalances"`
	} `json:"result"`
}

type getBlockResponse struct {
	Result struct {
		Timestamp string `json:"timestamp"`
	} `json:"result"`
}

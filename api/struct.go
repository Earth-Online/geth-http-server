package api

type TransactionObject struct {
	Hash             string `json:"hash,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	BlockHash        string `json:"blockHash,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`

	From     string `json:"from"`
	To       string `json:"to,omitempty"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
	Input    string `json:"input,omitempty"`
}

type TransactionInfo struct {
	From string `json:"from"`
	To   string `json:"to,omitempty"`
	GasUse string `json:"gasUsed"`
	Hash   string `json:"hash,omitempty"`
	Value    string `json:"value,omitempty"`

}


type PriceObject struct {
	Ethbtc string `json:"ethbtc"`
	Ethbtc_timestamp string  `json:"ethbtc_timestamp"`
	Ethusd string `json:"ethusd"`
	Ethusd_timestamp string   `json:"ethusd_timestamp s"`
}


package swap

type CryptoSwapReq struct {
	Reference  string      `json:"reference,omitempty"`
	Amount     int         `json:"amount"`
	FromCoin   string      `json:"fromCoin"`
	ToCoin     string      `json:"toCoin"`
	Blockchain string      `json:"blockchain"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

type GetCryptoSwapAmountOutReq struct {
	Amount     int    `json:"amount"`
	FromCoin   string `json:"fromCoin"`
	ToCoin     string `json:"toCoin"`
	Blockchain string `json:"blockchain"`
}

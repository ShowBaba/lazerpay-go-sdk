package transfer

type TransaferCryptoReq struct {
	Reference  string      `json:"reference,omitempty"`
	Amount     int         `json:"amount"`
	Recipient  string      `json:"recipient"`
	Coin       string      `json:"coin"`
	Blockchain string      `json:"blockchain"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

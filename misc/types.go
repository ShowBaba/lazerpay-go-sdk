package misc

type GetWalletBalanceReq struct {
	Coin string `json:"coin"`
}

type GetRateReq struct {
	Coin     string `json:"coin"`
	Currency string `json:"currency"`
}

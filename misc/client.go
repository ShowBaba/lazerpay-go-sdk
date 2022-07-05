package misc

import (
	"encoding/json"

	"github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	GetAcceptedCoins() (res *lazerpay.CustomResp, err error)
	GetWalletBalance(arg *GetWalletBalanceReq) (res *lazerpay.CustomResp, err error)
	GetRate(arg *GetRateReq) (res *lazerpay.CustomResp, err error) 
}

type apiImpl lazerpay.Context

func New(c lazerpay.Config) Client {
	ctx := apiImpl(lazerpay.NewContext(c))
	return &ctx
}

func (p *apiImpl) GetAcceptedCoins() (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "GET",
		Route:  []string{"misc", "get-accepted-coin"},
		Auth:   "PUB_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) GetWalletBalance(arg *GetWalletBalanceReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method:  "GET",
		Route:   []string{"misc", "get-wallet-balance"},
		Queries: map[string]string{"coin": arg.Coin},
		Auth:    "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) GetRate(arg *GetRateReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method:  "GET",
		Route:   []string{"misc", "get-rate"},
		Queries: map[string]string{"coin": arg.Coin, "currency": arg.Currency},
		Auth:    "PUB_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

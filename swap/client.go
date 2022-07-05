package swap

import (
	"encoding/json"
	"log"

	"github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	CryptoSwap(arg *CryptoSwapReq) (res *lazerpay.CustomResp, err error)
	GetCryptoSwapAmountOut(arg *GetCryptoSwapAmountOutReq) (res *lazerpay.CustomResp, err error)
}

type apiImpl lazerpay.Context

func New(c lazerpay.Config) Client {
	ctx := apiImpl(lazerpay.NewContext(c))
	return &ctx
}

func (p *apiImpl) CryptoSwap(arg *CryptoSwapReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "POST",
		Route:  []string{"swap", "crypto"},
		Auth:   "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, arg)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) GetCryptoSwapAmountOut(arg *GetCryptoSwapAmountOutReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "POST",
		Route:  []string{"swap", "crypto-amount-out"},
		Auth:   "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, arg)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

package transfer

import (
	"encoding/json"
	"log"

	"github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	TransferCrypto(arg *TransaferCryptoReq) (res *lazerpay.CustomResp, err error)
}

type apiImpl lazerpay.Context

func New(c lazerpay.Config) Client {
	ctx := apiImpl(lazerpay.NewContext(c))
	return &ctx
}

func (p *apiImpl) TransferCrypto(arg *TransaferCryptoReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "POST",
		Route:  []string{"transfer", "transfer-crypto"},
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

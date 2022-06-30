package payment

import (
	"encoding/json"
	"log"

	lazerpay "github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	InitializePayment(arg *InitPaymentReq) (res *lazerpay.CustomResp, err error)
	VerifyPayment(arg *VerifyPaymentReq) (res *lazerpay.CustomResp, err error)
}

type apiImpl lazerpay.Context

func New(c lazerpay.Config) Client {
	ctx := apiImpl(lazerpay.NewContext(c))
	return &ctx
}

func (p *apiImpl) InitializePayment(arg *InitPaymentReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "POST",
		Route:  []string{"payment", "create"},
		Auth:   "PUB_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, arg)
	if err != nil {
		log.Println("err: ", err)
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) VerifyPayment(arg *VerifyPaymentReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "GET",
		Route:  []string{"payment", "verify"},
		Auth:   "PUB_KEY",
		Identifier: arg.Identifier,
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}



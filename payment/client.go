package payment

import (
	"encoding/json"
	"fmt"
	"log"

	lazerpay "github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	InitializePayment(arg *InitPaymentReq) (res *InitPaymentResp, err error)
	VerifyPayment(id string) (res *VerifyPaymentResp, err error)
}

type apiImpl lazerpay.Context

var BASE_URL string

func (p *apiImpl) InitializePayment(arg *InitPaymentReq) (res *InitPaymentResp, err error) {
	url := BASE_URL + (*lazerpay.Context)(p).GetEndpoint("payment", "create")
	resp, err := (*lazerpay.Context)(p).SendRequest("POST", url, arg, nil, "PUB_KEY")
	if err != nil {
		log.Println("err: ", err)
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) VerifyPayment(id string) (res *VerifyPaymentResp, err error) {
	url := (*lazerpay.Context)(p).GetEndpoint("payment", "verify")
	urlWithParam := fmt.Sprintf(`%s/%s`, url, id)
	resp, err := (*lazerpay.Context)(p).SendRequest("GET", urlWithParam, nil, nil, "PUB_KEY")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func New(c lazerpay.Config) Client {
	ctx := apiImpl(lazerpay.NewContext(c))
	BASE_URL = (*lazerpay.Context)(&ctx).GetBaseURL(c)
	return &ctx
}

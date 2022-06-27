package payment

import (
	"encoding/json"
	"fmt"

	lazerpay "github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	InitializePayment(arg *InitPaymentReq) (res *InitPaymentResp, err error)
	VerifyPayment(id string) (res *VerifyPaymentResp, err error)
}

type apiImpl lazerpay.Context

func (p *apiImpl) InitializePayment(arg *InitPaymentReq) (res *InitPaymentResp, err error) {
	url := (*lazerpay.Context)(p).GetEndpoint("payment", "create")
	resp, err := (*lazerpay.Context)(p).SendRequest("POST", url, arg, nil, "PUB_KEY")
	if err != nil {
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
	return &ctx
}

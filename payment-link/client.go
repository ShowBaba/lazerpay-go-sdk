package paymentlink

import (
	"encoding/json"

	"github.com/ShowBaba/lazerpay-go-sdk"
)

type Client interface {
	CreatePaymentLink(arg *CreatePaymentLinkReq) (res *lazerpay.CustomResp, err error)
	UpdatePaymentLink(arg *UpdatePaymentLinkReq) (res *lazerpay.CustomResp, err error)
	GetAllPaymentLinks() (res *lazerpay.CustomResp, err error)
	GetPaymentLink(arg *GetPaymentLinkReq) (res *lazerpay.CustomResp, err error)
}

type apiImpl lazerpay.Context

func New(c lazerpay.Config) Client {
	ctx := apiImpl(lazerpay.NewContext(c))
	return &ctx
}

func (p *apiImpl) CreatePaymentLink(arg *CreatePaymentLinkReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "POST",
		Route:  []string{"payment_link", "create"},
		Auth:   "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, arg)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) UpdatePaymentLink(arg *UpdatePaymentLinkReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method:     "PUT",
		Route:      []string{"payment_link", "update"},
		Identifier: arg.Identifier,
		Auth:       "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, arg)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) GetAllPaymentLinks() (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method: "GET",
		Route:  []string{"payment_link", "fetch"},
		Auth:   "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

func (p *apiImpl) GetPaymentLink(arg *GetPaymentLinkReq) (res *lazerpay.CustomResp, err error) {
	reqData := lazerpay.Request{
		Method:     "GET",
		Route:      []string{"payment_link", "fetch"},
		Identifier: arg.Identifier,
		Auth:       "SEC_KEY",
	}
	resp, err := (*lazerpay.Context)(p).SendRequest(reqData, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}

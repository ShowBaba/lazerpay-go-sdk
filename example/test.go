package test

import (
	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/payment"
)

var (
	LAZER_PUBLIC_KEY = ""
	LAZER_SECRET_KEY = ""
)

func main() {
	config := lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
		Live:      true,
	}

	p := payment.New(config)
}

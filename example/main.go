package main

import (
	"fmt"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/payment"
)

var (
	LAZER_PUBLIC_KEY = "test-pub-key"
	LAZER_SECRET_KEY = "test-sec-key"
)

func main() {
	config := lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
		Live:      true,
	}

	p := payment.New(config)
	fmt.Println(payment.BASE_URL)
	args := payment.NewInitPaymentReq("YOUR_REFERENCE", "Samuel Shoyemi", "samwise858@gmail.com", "USDT", "USD", 100, true, map[string]string{"type": "Wallet fund"})
	resp, err := p.InitializePayment(args)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v", resp)
}

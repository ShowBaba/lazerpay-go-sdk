package lazerpay_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/payment"
)

// go test -v

var (
	LAZER_PUBLIC_KEY = "test-pub-key"
	LAZER_SECRET_KEY = "test-sec-key"
	ctx              lazerpay.Context
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
		Live:      true,
	}
)

func setup() {
	ctx = lazerpay.NewContext(config)
}

func TestRequests(t *testing.T) {
	setup()

	t.Run("GET Request", func(t *testing.T) {
		p := payment.New(config)
		args := payment.NewInitPaymentReq("YOUR_REFERENCE", "Samuel Shoyemi", "samwise858@gmail.com", "USDT", "USD", 100, true, map[string]string{"type": "Wallet fund"})
		resp, err := p.InitializePayment(args)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp)
	})

	t.Run("POST Request", func(t *testing.T) {

	})
}

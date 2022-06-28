package payment_test

import (
	"fmt"
	"testing"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/payment"
)

var (
	LAZER_PUBLIC_KEY = "test-pub-key"
	LAZER_SECRET_KEY = "test-sec-key"
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
		Live:      true,
	}
	client payment.Client
)

func setup() {
	client = payment.New(config)
}

func TestPayment(t *testing.T) {
	setup()

	t.Run("Initialize Payment", func(t *testing.T) {
		args := payment.NewInitPaymentReq("YOUR_REFERENCE", "Samuel Shoyemi", "samwise858@gmail.com", "USDT", "USD", 100, true, map[string]string{"type": "Wallet fund"})
		resp, err := client.InitializePayment(args)
		if err != nil {
			fmt.Printf(`error: %v`, err)
		}
		fmt.Printf("response: %v", resp)
	})

	t.Run("Verify Payment", func(t *testing.T) {
		resp, err := client.VerifyPayment("12334556")
		if err != nil {
			fmt.Printf(`error: %v`, err)
		}
		fmt.Printf("response: %v", resp)
	})
}

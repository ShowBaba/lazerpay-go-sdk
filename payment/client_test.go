package payment_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/payment"
)

var (
	LAZER_PUBLIC_KEY = "pk_test_cF00yS979TU2NguXHHjqLT05yknbaue7uqv7LwFD0alUEUvIue"
	LAZER_SECRET_KEY = "sk_test_Un4pKbE7Z9duLdQiBBQJIlTI7npLmd7PZfgUYQWjRJ5wNQz68y"
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
	}
	client payment.Client
	uniqueID string
)

func setup() {
	client = payment.New(config)
	uniqueID = uuid.New().String()
}

func TestPayment(t *testing.T) {
	setup()

	t.Run("Initialize Payment", func(t *testing.T) {
		arg := &payment.InitPaymentReq{
			Reference: uniqueID,
			CustomerName: "Samuel Shoyemi",
			CustomerEmail: "samwise858@gmail.com",
			Coin: "USDT",
			Currency: "USD",
			Amount: 100,
			AcceptPartialPayment: true,
			Metadata:  map[string]string{"type": "Wallet fund"},
		}
		resp, err := client.InitializePayment(arg)
		if err != nil {
			fmt.Printf(`error: %v`, err)
		}
		fmt.Printf("response: %v", resp)
	})

	t.Run("Verify Payment", func(t *testing.T) {
		arg := &payment.VerifyPaymentReq{
			Identifier: uniqueID,
		}
		resp, err := client.VerifyPayment(arg)
		if err != nil {
			fmt.Printf(`error: %v`, err)
		}
		fmt.Printf("response: %v", resp)
	})
}

package paymentlink_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/ShowBaba/lazerpay-go-sdk"
	paymentlink "github.com/ShowBaba/lazerpay-go-sdk/payment-link"
)

var (
	LAZER_PUBLIC_KEY = "pk_test_cF00yS979TU2NguXHHjqLT05yknbaue7uqv7LwFD0alUEUvIue"
	LAZER_SECRET_KEY = "sk_test_Un4pKbE7Z9duLdQiBBQJIlTI7npLmd7PZfgUYQWjRJ5wNQz68y"
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
	}
	client   paymentlink.Client
	uniqueID string
)

func setup() {
	client = paymentlink.New(config)
	uniqueID = uuid.New().String()
}

func TestPayment(t *testing.T) {
	setup()

	t.Run("Create PaymentLink", func(t *testing.T) {
		arg := &paymentlink.CreatePaymentLinkReq{
			Title:       "test link",
			Description: "lorem ipsum",
			Logo:        "https://assets.audiomack.com/fireboydml/bbbd8710eff038d4f603cc39ec94a6a6c2c5b6f4100b28d62557d10d87246f27.jpeg?width=340&height=340&max=true",
			RedirectURL: "",
			Amount:      100,
			Currency:    "USD",
			Type:        "standard",
		}
		resp, err := client.CreatePaymentLink(arg)
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v\n\n", resp)
	})

	t.Run("Update Payment Link", func(t *testing.T) {
		arg := &paymentlink.UpdatePaymentLinkReq{
			Status:     "inactive",
			Identifier: uniqueID,
		}
		resp, err := client.UpdatePaymentLink(arg)
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v\n\n", resp)
	})

	t.Run("Get All Payment Links", func(t *testing.T) {
		resp, err := client.GetAllPaymentLinks()
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v\n\n", resp)
	})

	t.Run("Get One Payment Link", func(t *testing.T) {
		arg := &paymentlink.GetPaymentLinkReq{
			Identifier: uniqueID,
		}
		resp, err := client.GetPaymentLink(arg)
		if err != nil {
			fmt.Printf(`error: %v\n\n`, err)
		}
		fmt.Printf("response: %v\n\n", resp)
	})
}

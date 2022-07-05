package misc_test

import (
	"fmt"
	"testing"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/misc"
	"github.com/google/uuid"
)

var (
	LAZER_PUBLIC_KEY = "pk_test_cF00yS979TU2NguXHHjqLT05yknbaue7uqv7LwFD0alUEUvIue"
	LAZER_SECRET_KEY = "sk_test_Un4pKbE7Z9duLdQiBBQJIlTI7npLmd7PZfgUYQWjRJ5wNQz68y"
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
	}
	client   misc.Client
	uniqueID string
)

func setup() {
	client = misc.New(config)
	uniqueID = uuid.New().String()
}

func TestMisc(t *testing.T) {
	setup()

	t.Run("Get All Accepted Coins", func(t *testing.T) {
		resp, err := client.GetAcceptedCoins()
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v", resp)
	})

	t.Run("Get Wallet Balance", func(t *testing.T) {
		arg := &misc.GetWalletBalanceReq{
			Coin: "USDT",
		}
		resp, err := client.GetWalletBalance(arg)
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v", resp)
	})

	t.Run("Get Rate", func(t *testing.T) {
		arg := &misc.GetRateReq{
			Coin: "USDT",
			Currency: "USDT",
		}
		resp, err := client.GetRate(arg)
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v", resp)
	})
}

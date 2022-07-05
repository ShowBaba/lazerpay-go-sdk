package swap_test

import (
	"fmt"
	"testing"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/swap"
	"github.com/google/uuid"
)

var (
	LAZER_PUBLIC_KEY = "pk_test_cF00yS979TU2NguXHHjqLT05yknbaue7uqv7LwFD0alUEUvIue"
	LAZER_SECRET_KEY = "sk_test_Un4pKbE7Z9duLdQiBBQJIlTI7npLmd7PZfgUYQWjRJ5wNQz68y"
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
	}
	client   swap.Client
	uniqueID string
)

func setup() {
	client = swap.New(config)
	uniqueID = uuid.New().String()
}

func TestSwap(t *testing.T) {
	setup()

	t.Run("Crypto Swap", func(t *testing.T) {
		arg := &swap.CryptoSwapReq{
			Reference:  uniqueID,
			Amount:     100,
			FromCoin:   "BUSD",
			ToCoin:     "USDT",
			Blockchain: "Binance Smart Chain",
		}
		resp, err := client.CryptoSwap(arg)
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v\n\n", resp)
	})

	t.Run("Get Crypto Swap Amount Out", func(t *testing.T) {
		arg := &swap.GetCryptoSwapAmountOutReq{
			Amount:     100,
			FromCoin:   "BUSD",
			ToCoin:     "USDT",
			Blockchain: "Binance Smart Chain",
		}
		resp, err := client.GetCryptoSwapAmountOut(arg)
		if err != nil {
			t.Errorf("unexpected error occured; err: %v", err)
			return
		}
		fmt.Printf("response: %v\n\n", resp)
	})
}

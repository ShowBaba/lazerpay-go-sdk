package transfer_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/ShowBaba/lazerpay-go-sdk"
	"github.com/ShowBaba/lazerpay-go-sdk/transfer"
)

var (
	LAZER_PUBLIC_KEY = "pk_test_cF00yS979TU2NguXHHjqLT05yknbaue7uqv7LwFD0alUEUvIue"
	LAZER_SECRET_KEY = "sk_test_Un4pKbE7Z9duLdQiBBQJIlTI7npLmd7PZfgUYQWjRJ5wNQz68y"
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
	}
	client   transfer.Client
	uniqueID string
)

func setup() {
	client = transfer.New(config)
	uniqueID = uuid.New().String()
}

func TestTransfer(t *testing.T) {
	setup()

	t.Run("Transfer Crypto", func(t *testing.T) {
		arg := &transfer.TransaferCryptoReq{
			Reference:  uniqueID,
			Amount:     100,
			Recipient:  "0x0B4d358D349809037003F96A3593ff9015E89efA",
			Coin:       "USDT",
			Blockchain: "Binance Smart Chain",
			Metadata:   map[string]string{"type": "Crypto transfer"},
		}
		resp, err := client.TransferCrypto(arg)
		if err != nil {
			fmt.Printf(`error: %v`, err)
		}
		fmt.Printf("response: %v", resp)
	})
}

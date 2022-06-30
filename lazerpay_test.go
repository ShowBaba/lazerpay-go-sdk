package lazerpay_test

import (
	"testing"

	"github.com/ShowBaba/lazerpay-go-sdk"
)

// go test -v

var (
	LAZER_PUBLIC_KEY = "test-pub-key"
	LAZER_SECRET_KEY = "test-sec-key"
	ctx              lazerpay.Context
	config           = lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
	}
)

func setup() {
	ctx = lazerpay.NewContext(config)
}

func TestRequests(t *testing.T) {
	setup()

	t.Run("GET Request", func(t *testing.T) {
		
	})

	t.Run("POST Request", func(t *testing.T) {

	})
}

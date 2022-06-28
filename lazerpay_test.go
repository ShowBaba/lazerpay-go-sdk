package lazerpay_test

import (
	"testing"

	"github.com/ShowBaba/lazerpay-go-sdk"
)

var (
	LAZER_PUBLIC_KEY = "test-pub-key"
	LAZER_SECRET_KEY = "test-sec-key"
	ctx lazerpay.Context
)

func setup()  {
	config := lazerpay.Config{
		ApiPubKey: LAZER_PUBLIC_KEY,
		ApiSecKey: LAZER_SECRET_KEY,
		Live:      true,
	}
	ctx = lazerpay.NewContext(config)
}

func TestRequests(t *testing.T){
	setup()

	t.Run("GET request", func(t *testing.T){

	})

	t.Run("POST request", func(t *testing.T) {
		
	})
}

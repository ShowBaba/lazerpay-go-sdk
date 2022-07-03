# Lazerpay SDK for Go [UNOFFICIAL][Under Construction]

## Installation
```sh
go get github.com/ShowBaba/lazerpay-go-sdk
```

This SDK is built so you can import relavant namespace(s) only.

## Exported sub-packages

* `go get github.com/ShowBaba/lazerpay-go-sdk/payment`
* `go get github.com/ShowBaba/lazerpay-go-sdk/payment-link`
* `go get github.com/ShowBaba/lazerpay-go-sdk/transfer`

With the base at:
* `go get github.com/ShowBaba/lazerpay-go-sdk`

## Usage 

### Configuration
```go
  import "github.com/ShowBaba/lazerpay-go-sdk"

  func main() {
    config := lazerpay.Config{
      apiPubKey: LAZER_PUBLIC_KEY,
      apiSecKey: LAZER_SECRET_KEY,
      Live: true,
    }
  }
```

### Making API calls

## Payment

#### `Initialize Payment`
This describes to allow your customers to initiate a crypto payment transfer.

```go
  import "github.com/ShowBaba/lazerpay-go-sdk/payment"

  client := payment.New(config)

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
```
#### `Confirm Payment`
This describes to allow you confirm your customers transaction after payment has been made.

```go
  arg := &payment.VerifyPaymentReq{
    Identifier: uniqueID,
  }
  resp, err := client.VerifyPayment(arg)
  if err != nil {
    fmt.Printf(`error: %v`, err)
  }
  fmt.Printf("response: %v", resp)
```

## Payment Links

#### `Create a payment link`
This describes to allow you create a Payment link programatically

```go
  import (
	  paymentlink "github.com/ShowBaba/lazerpay-go-sdk/payment-link"
  )
  
	client = paymentlink.New(config)

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
    fmt.Printf(`error: %v\n`, err)
  }
  fmt.Printf("response: %v\n\n", resp)
```
#### `Update a payment link`
This describes disabling or enabling a payment link by updating it

```go
  arg := &paymentlink.UpdatePaymentLinkReq{
    Status:     "inactive",
    Identifier: uniqueID,
  }
  resp, err := client.UpdatePaymentLink(arg)
  if err != nil {
    fmt.Printf(`error: %v\n`, err)
  }
  fmt.Printf("response: %v\n\n", resp)
```

#### `Get all payment links`
This describes to allow you get all Payment links created

```go
  resp, err := client.GetAllPaymentLinks()
  if err != nil {
    fmt.Printf(`error: %v\n`, err)
  }
  fmt.Printf("response: %v\n\n", resp)
```

#### `Get a single payment link`
This describes to allow you get a Payment link by it's identifier

```go
	arg := &paymentlink.GetPaymentLinkReq{
			Identifier: uniqueID,
  }
  resp, err := client.GetPaymentLink(arg)
  if err != nil {
    fmt.Printf(`error: %v\n\n`, err)
  }
  fmt.Printf("response: %v\n\n", resp)
```

#### `Crypto Payout`
This describes to allow you withdraw the crypto in their lazerpay balance to an external address

```go
	import "github.com/ShowBaba/lazerpay-go-sdk/transfer"

	client = transfer.New(config)

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
```
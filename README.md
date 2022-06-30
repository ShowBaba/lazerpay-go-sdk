# Lazerpay SDK for Go [UNOFFICIAL][Under Construction]

## Installation
```sh
go get github.com/ShowBaba/lazerpay-go-sdk
```

This SDK is built so you can import relavant namespace(s) only.

The following sub-packages are exported:

* `go get github.com/ShowBaba/lazerpay-go-sdk/payment`

With the base at:
* `go get github.com/ShowBaba/lazerpay-go-sdk`

## Usage

```go
  import "github.com/ShowBaba/lazerpay-go-sdk"
  import "github.com/ShowBaba/lazerpay-go-sdk/payment"

  func main() {
    config := lazerpay.Config{
      apiPubKey: LAZER_PUBLIC_KEY,
      apiSecKey: LAZER_SECRET_KEY,
      Live: true,
    }

    client := payment.New(config)
  }
```

### Making API calls

```go
  // initialize payment
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

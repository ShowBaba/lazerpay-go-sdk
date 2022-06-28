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

    p := payment.New(config)
  }
```

### Making API calls

```go
  // initialize payment
  args := payment.NewInitPaymentReq("YOUR_REFERENCE","Samuel Shoyemi", "samwise858@gmail.com", "USDT", "USD", 100, true, map[string]string{"type": "Wallet fund"})

   resp, err := p.InitializePayment(args); if err != nil {
    	fmt.Printf("error: %v", err)
  }
  fmt.Printf("response: %v", resp)
```

package payment

type InitPaymentReq struct {
	Reference            string      `json:"reference,omitempty"`
	Amount               int      `json:"amount"`
	CustomerName         string      `json:"customer_name"`
	CustomerEmail        string      `json:"customer_email"`
	Coin                 string      `json:"coin"`
	Currency             string      `json:"currency"`
	AcceptPartialPayment bool        `json:"accept_partial_payment,omitempty"`
	Metadata             interface{} `json:"metadata,omitempty"`
}

func NewInitPaymentReq(Reference, CustomerName, CustomerEmail, Coin, Currency string, Amount int, AcceptPartialPayment bool, Metadata interface{}) *InitPaymentReq {
	c := new(InitPaymentReq)
	c.Reference = Reference
	c.Amount = Amount
	c.CustomerName = CustomerName
	c.CustomerEmail = CustomerEmail
	c.Coin = Coin
	c.Currency = Currency
	c.AcceptPartialPayment = AcceptPartialPayment
	c.Metadata = Metadata
	return c
}

type InitPaymentResp struct{}

type VerifyPaymentResp struct{}

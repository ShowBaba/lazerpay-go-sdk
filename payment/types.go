package payment

type InitPaymentReq struct {
	Reference string `json:"reference,omitempty"`
	Amount string `json:"amount"`
	CustomerName string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	Coin string `json:"coin"`
	Currency string `json:"currency"`
	AcceptPartialPayment bool `json:"accept_partial_payment,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

type InitPaymentResp struct {}

type VerifyPaymentResp struct{}
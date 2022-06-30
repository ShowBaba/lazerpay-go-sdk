package payment

type InitPaymentReq struct {
	Reference            string      `json:"reference,omitempty"`
	CustomerName         string      `json:"customer_name"`
	CustomerEmail        string      `json:"customer_email"`
	Coin                 string      `json:"coin"`
	Currency             string      `json:"currency"`
	Amount               int         `json:"amount"`
	AcceptPartialPayment bool        `json:"accept_partial_payment,omitempty"`
	Metadata             interface{} `json:"metadata,omitempty"`
}

type VerifyPaymentReq struct {
	Identifier string `json:"identifier"`
}


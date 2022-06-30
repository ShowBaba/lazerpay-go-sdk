package paymentlink

type CreatePaymentLinkReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Type        string `json:"type,omitempty"`
}

type UpdatePaymentLinkReq struct {
	Status     string `json:"status"`
	Identifier string `json:"identifier"`
}

type GetPaymentLinkReq struct {
	Identifier string `json:"identifier"`
}

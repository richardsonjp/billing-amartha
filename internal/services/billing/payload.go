package billing

type BillingCreatePayload struct {
	LoanID string `json:"loan_id"`
	Amount int64  `json:"amount"`
}

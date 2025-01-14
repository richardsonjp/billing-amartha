package billing

type BillingPaymentPayload struct {
	LoanID string `json:"loan_id" validate:"required"`
}

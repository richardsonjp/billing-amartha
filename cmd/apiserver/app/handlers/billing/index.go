package billing

import (
	"go-skeleton/internal/services"
)

type BillingHandler struct {
	billingService service.BillingService
}

func NewBillingHandler(billingService service.BillingService) *BillingHandler {
	return &BillingHandler{
		billingService: billingService,
	}
}

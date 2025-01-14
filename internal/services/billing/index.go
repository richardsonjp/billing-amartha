package billing

import (
	"context"
	"go-skeleton/internal/model"
	repos "go-skeleton/internal/repositories"
)

type BillingService interface {
	CreateBilling(ctx context.Context, payload BillingCreatePayload) error
	GetUnpaidBillings(ctx context.Context, loanID string) ([]model.Billing, error)
	UpdateBillingAsPaid(ctx context.Context, billing model.Billing) error
	DoPaymentBilling(ctx context.Context, loanID string) error
}

type billingService struct {
	billingRepo repos.BillingRepo
}

func NewBillingService(billingRepo repos.BillingRepo) BillingService {
	return &billingService{
		billingRepo: billingRepo,
	}
}

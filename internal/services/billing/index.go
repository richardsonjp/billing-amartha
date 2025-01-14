package billing

import (
	"context"
	"go-skeleton/internal/model"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services/borrower"
	"go-skeleton/internal/services/loan"
)

type BillingService interface {
	CreateBilling(ctx context.Context, payload BillingCreatePayload) error
	GetBillings(ctx context.Context, loanID string) ([]model.Billing, error)
	GetUnpaidBillings(ctx context.Context, loanID string) ([]model.Billing, error)
	UpdateBillingAsPaid(ctx context.Context, billing model.Billing) error
	DoPaymentBilling(ctx context.Context, loanID string) error
}

type billingService struct {
	billingRepo     repos.BillingRepo
	borrowerService borrower.BorrowerService
	loanService     loan.LoanService
}

func NewBillingService(billingRepo repos.BillingRepo, borrowerService borrower.BorrowerService, loanService loan.LoanService) BillingService {
	return &billingService{
		billingRepo:     billingRepo,
		borrowerService: borrowerService,
		loanService:     loanService,
	}
}

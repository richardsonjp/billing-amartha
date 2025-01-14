package billing

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/internal/services/borrower"
)

func (s *billingService) CreateBilling(ctx context.Context, payload BillingCreatePayload) error {
	s.billingRepo.Create(ctx, &model.Billing{
		LoanID: payload.LoanID,
		Amount: payload.Amount,
		Status: "outstanding",
	})
	return nil
}

func (s *billingService) GetBillings(ctx context.Context, loanID string) ([]model.Billing, error) {
	return s.billingRepo.GetByLoanID(ctx, loanID)
}

func (s *billingService) GetUnpaidBillings(ctx context.Context, loanID string) ([]model.Billing, error) {
	return s.billingRepo.GetUnpaidBillings(ctx, loanID)
}

func (s *billingService) UpdateBillingAsPaid(ctx context.Context, billing model.Billing) error {
	billing.Status = "paid"

	_, err := s.billingRepo.Update(ctx, billing, "status")
	return err
}

func (s *billingService) DoPaymentBilling(ctx context.Context, loanID string) error {
	loanData, err := s.loanService.GetByID(ctx, loanID)
	if err != nil {
		return err
	}
	billings, err := s.GetUnpaidBillings(ctx, loanData.ID)
	if err != nil {
		return err
	}
	for _, billing := range billings {
		err = s.UpdateBillingAsPaid(ctx, billing)
		if err != nil {
			return err
		}
	}

	return s.borrowerService.ToggleBorrowerStatus(ctx, borrower.BorrowerUpdatePayload{BorrowerID: loanData.BorrowerID, Status: "normal"})
}

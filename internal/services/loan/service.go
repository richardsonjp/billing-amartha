package loan

import (
	"context"
	"go-skeleton/internal/model"
	"time"
)

func (s *loanService) GetByID(ctx context.Context, loanID string) (*model.Loan, error) {
	return s.loanRepo.GetOne(ctx, &model.Loan{ID: loanID})
}

func (s *loanService) GetLoanByDayRange(ctx context.Context) ([]model.Loan, error) {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	loansOverdue, err := s.loanRepo.GetLoansWithoutRecentBilling(ctx, sevenDaysAgo)
	if err != nil {
		return nil, err
	}
	return loansOverdue, nil
}

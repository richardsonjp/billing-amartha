package loan

import (
	"context"
	"go-skeleton/internal/model"
	repos "go-skeleton/internal/repositories"
)

type LoanService interface {
	GetByID(ctx context.Context, loanID string) (*model.Loan, error)
	GetLoanByDayRange(ctx context.Context) ([]model.Loan, error)
}

type loanService struct {
	loanRepo repos.LoanRepo
}

func NewLoanService(loanRepo repos.LoanRepo) LoanService {
	return &loanService{
		loanRepo: loanRepo,
	}
}

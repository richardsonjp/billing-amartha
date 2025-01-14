package loan

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
	"time"
)

type LoanRepo interface {
	Create(ctx context.Context, m *model.Loan) (*model.Loan, error)
	GetOne(ctx context.Context, m *model.Loan) (*model.Loan, error)
	GetLoansWithoutRecentBilling(ctx context.Context, timeRange time.Time) ([]model.Loan, error)
}

type loanRepo struct {
	dbdget db.DBGormDelegate
}

func NewLoanRepo(dbdget db.DBGormDelegate) LoanRepo {
	return &loanRepo{
		dbdget: dbdget,
	}
}

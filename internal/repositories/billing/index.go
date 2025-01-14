package billing

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type BillingRepo interface {
	Create(ctx context.Context, m *model.Billing) (*model.Billing, error)
	Update(ctx context.Context, m model.Billing, updatedFields ...string) (int64, error)
	GetByLoanID(ctx context.Context, loanID string) ([]model.Billing, error)
	GetUnpaidBillings(ctx context.Context, loanID string) ([]model.Billing, error)
}

type billingRepo struct {
	dbdget db.DBGormDelegate
}

func NewBillingRepo(dbdget db.DBGormDelegate) BillingRepo {
	return &billingRepo{
		dbdget: dbdget,
	}
}

package loan

import (
	"context"
	"go-skeleton/internal/model"
	"time"
)

func (r *loanRepo) Create(ctx context.Context, m *model.Loan) (*model.Loan, error) {
	if err := r.dbdget.Get(ctx).
		Create(m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *loanRepo) GetOne(ctx context.Context, m *model.Loan) (*model.Loan, error) {
	query := r.dbdget.Get(ctx).Where(m)

	if err := query.Last(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *loanRepo) GetLoansWithoutRecentBilling(ctx context.Context, timeRange time.Time) ([]model.Loan, error) {
	var loans []model.Loan

	if err := r.dbdget.Get(ctx).
		Model(&model.Loan{}).
		Joins("LEFT JOIN billing ON billing.loan_id = billing.id AND billing.created_at >= ?", timeRange).
		Where("billing.id IS NULL").
		Find(&loans).Error; err != nil {
		return nil, err
	}

	return loans, nil
}

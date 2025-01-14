package billing

import (
	"context"
	"go-skeleton/internal/model"
)

func (r *billingRepo) Create(ctx context.Context, m *model.Billing) (*model.Billing, error) {
	if err := r.dbdget.Get(ctx).
		Create(m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *billingRepo) Update(ctx context.Context, m model.Billing, updatedFields ...string) (int64, error) {
	query := r.dbdget.Get(ctx)
	if len(updatedFields) > 0 {
		updatedFields = append(updatedFields, "updated_at")
		query = query.Select(updatedFields)
	}
	query.Updates(m)

	// execute query
	if err := query.Find(&m).Error; err != nil {
		return 0, err
	}
	return query.RowsAffected, nil
}

func (r *billingRepo) GetByLoanID(ctx context.Context, loanID string) ([]model.Billing, error) {
	var billings []model.Billing

	query := r.dbdget.Get(ctx).Where("loan_id = ?", loanID)

	if err := query.Find(&billings).Error; err != nil {
		return nil, err
	}

	return billings, nil
}

func (r *billingRepo) GetUnpaidBillings(ctx context.Context, loanID string) ([]model.Billing, error) {
	var billings []model.Billing

	query := r.dbdget.Get(ctx).Where("loan_id = ?", loanID).Where("status = ?", "outstanding")

	if err := query.Find(&billings).Error; err != nil {
		return nil, err
	}

	return billings, nil
}

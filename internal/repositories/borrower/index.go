package borrower

import (
	"context"
	"go-skeleton/internal/model"
	"go-skeleton/pkg/clients/db"
)

type BorrowerRepo interface {
	Create(ctx context.Context, m *model.Borrower) (*model.Borrower, error)
	GetOne(ctx context.Context, m *model.Borrower) (*model.Borrower, error)
	Update(ctx context.Context, param *model.Borrower, updatedFields ...string) (int64, error)
}

type borrowerRepo struct {
	dbdget db.DBGormDelegate
}

func NewBorrowerRepo(dbdget db.DBGormDelegate) BorrowerRepo {
	return &borrowerRepo{
		dbdget: dbdget,
	}
}

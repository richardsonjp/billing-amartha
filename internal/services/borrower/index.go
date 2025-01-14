package borrower

import (
	"context"
	"go-skeleton/internal/model"
	repos "go-skeleton/internal/repositories"
)

type BorrowerService interface {
	ToggleBorrowerStatus(ctx context.Context, payload BorrowerUpdatePayload) error
	GetBorrower(ctx context.Context, borrowerID string) (*model.Borrower, error)
}

type borrowerService struct {
	borrowerRepo repos.BorrowerRepo
}

func NewBorrowerService(borrowerRepo repos.BorrowerRepo) BorrowerService {
	return &borrowerService{
		borrowerRepo: borrowerRepo,
	}
}

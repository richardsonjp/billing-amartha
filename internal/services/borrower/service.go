package borrower

import (
	"context"
	"go-skeleton/internal/model"
)

func (s *borrowerService) ToggleBorrowerStatus(ctx context.Context, payload BorrowerUpdatePayload) error {
	data, err := s.GetBorrower(ctx, payload.BorrowerID)
	if err != nil {
		return err
	}

	data.Status = payload.Status

	_, err = s.borrowerRepo.Update(ctx, data, "status")
	return err
}

func (s *borrowerService) GetBorrower(ctx context.Context, borrowerID string) (*model.Borrower, error) {
	return s.borrowerRepo.GetOne(ctx, &model.Borrower{
		ID: borrowerID,
	})
}

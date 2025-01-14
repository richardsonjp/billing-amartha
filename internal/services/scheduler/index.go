package scheduler

import (
	"context"
	"go-skeleton/internal/services/billing"
	"go-skeleton/internal/services/borrower"
	"go-skeleton/internal/services/loan"
	"sync"
)

type SchedulerService interface {
	ToggleScheduler(ctx context.Context, toggleBool bool)
}

type schedulerService struct {
	borrowerService borrower.BorrowerService
	loanService     loan.LoanService
	billingService  billing.BillingService

	isRunning bool               // Tracks whether the scheduler is running
	mu        sync.Mutex         // Mutex to synchronize access to isRunning
	cancelFn  context.CancelFunc // Context cancel function to stop the scheduler

}

func NewSchedulerService(borrowerService borrower.BorrowerService, loanService loan.LoanService, billingService billing.BillingService) SchedulerService {
	return &schedulerService{
		borrowerService: borrowerService,
		loanService:     loanService,
		billingService:  billingService,
	}
}

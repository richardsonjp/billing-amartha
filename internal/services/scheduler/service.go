package scheduler

import (
	"context"
	"go-skeleton/internal/services/billing"
	"go-skeleton/internal/services/borrower"
	"log"
	"time"
)

func (s *schedulerService) StartDailyCheck(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Starting daily check service...")
			if err := s.GenerateBilling(ctx); err != nil {
				log.Println("Error during daily check service:", err)
			} else {
				log.Println("Finished daily check service...")
			}
		case <-ctx.Done():
			log.Println("Scheduler stopped")
			return
		}
	}
}

func (s *schedulerService) ToggleScheduler(ctx context.Context, toggleBool bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if toggleBool {
		if s.isRunning {
			log.Println("Scheduler is already running")
			return
		}
		// Start the scheduler in the background
		ctx, cancel := context.WithCancel(ctx)
		s.cancelFn = cancel
		s.isRunning = true
		go s.StartDailyCheck(ctx)
		log.Println("Scheduler started")
	} else {
		if !s.isRunning {
			log.Println("Scheduler is already stopped")
			return
		}
		// Stop the scheduler
		s.cancelFn()
		s.isRunning = false
		log.Println("Scheduler stopped")
	}
}

func (s *schedulerService) GenerateBilling(ctx context.Context) error {
	loans, err := s.loanService.GetLoanByDayRange(ctx)
	if err != nil {
		return err
	}

	for _, loan := range loans {
		err := s.billingService.CreateBilling(ctx, billing.BillingCreatePayload{
			LoanID: loan.ID,
			Amount: loan.Amount,
		})
		if err != nil {
			return err
		}
		bills, err := s.billingService.GetUnpaidBillings(ctx, loan.ID)
		if err != nil {
			return err
		}
		if len(bills) > 1 {
			err = s.borrowerService.ToggleBorrowerStatus(ctx, borrower.BorrowerUpdatePayload{
				BorrowerID: loan.BorrowerID,
				Status:     "delinquent",
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

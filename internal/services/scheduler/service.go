package scheduler

import (
	"context"
	"go-skeleton/internal/services/billing"
	"go-skeleton/internal/services/borrower"
	timeutil "go-skeleton/pkg/utils/time"
	"log"
	"time"
)

func (s *schedulerService) StartDailyCheck(ctx context.Context) {
	// Function to handle the daily check execution
	runDailyCheck := func() {
		log.Println("Starting daily check service...")
		if err := s.GenerateBilling(ctx); err != nil {
			log.Println("Error during daily check service:", err)
		} else {
			log.Println("Finished daily check service...")
		}
	}

	// Run immediately on start
	runDailyCheck()

	// Calculate the duration until midnight by location
	now := time.Now()
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	timeUntilNextRun := time.Until(nextMidnight)

	// Create a timer for the initial wait
	timer := time.NewTimer(timeUntilNextRun)

	for {
		select {
		case <-timer.C:
			// Run the daily check
			runDailyCheck()

			// Switch to using a ticker for subsequent runs every 24 hours
			timer.Reset(24 * time.Hour)

		case <-ctx.Done():
			log.Println("Scheduler stopped")
			timer.Stop()
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
		currentBills, err := s.billingService.GetBillings(ctx, loan.ID)
		if err != nil {
			return err
		}
		actualTotalWeeks := timeutil.WeekDifferenceCounter(timeutil.Now(), loan.CreatedAt)
		missedBillCounter := actualTotalWeeks - len(currentBills)
		for missedBillCounter > 0 {
			billingAmount := loan.TotalAmount / loan.TotalWeeks
			err = s.billingService.CreateBilling(ctx, billing.BillingCreatePayload{
				LoanID: loan.ID,
				Amount: billingAmount,
			})
			if err != nil {
				return err
			}
			missedBillCounter--
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

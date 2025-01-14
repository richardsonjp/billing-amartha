package service

import (
	"go-skeleton/internal/services/billing"
	"go-skeleton/internal/services/borrower"
	"go-skeleton/internal/services/loan"
	"go-skeleton/internal/services/scheduler"
)

// put handlers alias
type (
	BorrowerService  = borrower.BorrowerService
	LoanService      = loan.LoanService
	BillingService   = billing.BillingService
	SchedulerService = scheduler.SchedulerService
)

var (
	NewBorrowerService  = borrower.NewBorrowerService
	NewLoanService      = loan.NewLoanService
	NewBillingService   = billing.NewBillingService
	NewSchedulerService = scheduler.NewSchedulerService
)

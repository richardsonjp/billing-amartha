package repos

import (
	"go-skeleton/internal/repositories/billing"
	"go-skeleton/internal/repositories/borrower"
	"go-skeleton/internal/repositories/loan"
	"go-skeleton/internal/repositories/tx"
)

// put repos alias
type (
	BorrowerRepo = borrower.BorrowerRepo
	LoanRepo     = loan.LoanRepo
	BillingRepo  = billing.BillingRepo

	TxRepo = tx.TxRepo
)

var (
	NewBorrowerRepo = borrower.NewBorrowerRepo
	NewLoanRepo     = loan.NewLoanRepo
	NewBillingRepo  = billing.NewBillingRepo

	NewTxRepo = tx.NewTxRepo
)

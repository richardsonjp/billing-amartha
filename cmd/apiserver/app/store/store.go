package store

import (
	"go-skeleton/cmd/apiserver/app/handlers"
	"go-skeleton/config"
	repos "go-skeleton/internal/repositories"
	"go-skeleton/internal/services"
	"go-skeleton/pkg/clients/db"
)

var (
	// handlers
	BillingHandler   *handler.BillingHandler
	SchedulerHandler *handler.SchedulerHandler

	// services
	BorrowerService  service.BorrowerService
	LoanService      service.LoanService
	BillingService   service.BillingService
	SchedulerService service.SchedulerService

	// repos
	BorrowerRepo repos.BorrowerRepo
	LoanRepo     repos.LoanRepo
	BillingRepo  repos.BillingRepo

	TxRepo repos.TxRepo
)

// Init application global variable with single instance
func InitDI() {
	// setup resources
	dbdget := db.NewDBdelegate(config.Config.DB.Debug)
	dbdget.Init()

	// repos
	BorrowerRepo = repos.NewBorrowerRepo(dbdget)
	LoanRepo = repos.NewLoanRepo(dbdget)
	BillingRepo = repos.NewBillingRepo(dbdget)

	TxRepo = repos.NewTxRepo(dbdget)

	// servicesWalletRepo
	BorrowerService = service.NewBorrowerService(BorrowerRepo)
	LoanService = service.NewLoanService(LoanRepo)
	BillingService = service.NewBillingService(BillingRepo, BorrowerService, LoanService)
	SchedulerService = service.NewSchedulerService(BorrowerService, LoanService, BillingService)

	// handlers
	BillingHandler = handler.NewBillingHandler(BillingService)
	SchedulerHandler = handler.NewSchedulerHandler(SchedulerService)
}

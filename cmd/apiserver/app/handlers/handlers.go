package handler

import (
	"go-skeleton/cmd/apiserver/app/handlers/billing"
	"go-skeleton/cmd/apiserver/app/handlers/scheduler"
)

// put handlers alias
type (
	BillingHandler   = billing.BillingHandler
	SchedulerHandler = scheduler.SchedulerHandler
)

var (
	NewBillingHandler   = billing.NewBillingHandler
	NewSchedulerHandler = scheduler.NewSchedulerHandler
)

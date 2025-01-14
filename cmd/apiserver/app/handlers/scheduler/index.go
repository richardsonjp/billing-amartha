package scheduler

import (
	"go-skeleton/internal/services"
)

type SchedulerHandler struct {
	schedulerService service.SchedulerService
}

func NewSchedulerHandler(schedulerService service.SchedulerService) *SchedulerHandler {
	return &SchedulerHandler{
		schedulerService: schedulerService,
	}
}

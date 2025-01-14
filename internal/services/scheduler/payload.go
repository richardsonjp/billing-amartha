package scheduler

type ToggleSchedulerWorker struct {
	Toggle bool `json:"toggle" validate:"required"`
}

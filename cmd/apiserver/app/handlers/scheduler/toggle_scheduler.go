package scheduler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/services/scheduler"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *SchedulerHandler) ToggleScheduler(ctx *gin.Context) {
	param := scheduler.ToggleSchedulerWorker{}
	if err := ctx.ShouldBindWith(&param, binding.JSON); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(param); err != nil {
		errors.ErrorString(ctx, errorMessage)
		return
	}
	h.schedulerService.ToggleScheduler(ctx, param.Toggle)

	ctx.JSON(201, api.Message{
		Message: "success",
	})
}

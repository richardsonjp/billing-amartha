package billing

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-skeleton/internal/repositories/billing"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *BillingHandler) BillingPayment(ctx *gin.Context) {
	param := billing.BillingPaymentPayload{}
	if err := ctx.ShouldBindWith(&param, binding.JSON); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}
	if errorMessage, err := validator.Validate(param); err != nil {
		errors.ErrorString(ctx, errorMessage)
		return
	}

	err := h.billingService.DoPaymentBilling(ctx, param.LoanID)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(201, api.Message{
		Message: "success",
	})
}

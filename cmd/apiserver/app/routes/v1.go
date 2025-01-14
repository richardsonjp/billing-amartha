package routes

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/cmd/apiserver/app/store"
)

func V1Route(group *gin.RouterGroup) {
	group.POST("/toggle/scheduler", store.SchedulerHandler.ToggleScheduler)
	group.POST("/billing-payment", store.BillingHandler.BillingPayment)
}

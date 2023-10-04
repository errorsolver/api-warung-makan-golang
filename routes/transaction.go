package routes

import (
	"github.com/gin-gonic/gin"

	"golang-wm-api/middleware"
)

// type TransactionRoutes struct{}

func (RoutesCollection) TransactionRoutes(r *gin.RouterGroup) {
	r.GET("/transaction", middleware.AdminAuth(), CC.GetAllTransaction)
	r.GET("/transaction/:id", middleware.UserAuth, CC.GetTransactionByID)
	r.POST("/transaction", middleware.UserAuth, CC.CreateTransactionDetails)
}

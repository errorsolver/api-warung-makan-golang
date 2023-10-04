package routes

import (
	"github.com/gin-gonic/gin"
)

func (RoutesCollection) AdminRoutes(r *gin.RouterGroup) {
	r.POST("/login", CC.Login)
}
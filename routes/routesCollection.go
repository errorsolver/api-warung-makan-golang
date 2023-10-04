package routes

import (
	"golang-wm-api/controllers"

	"github.com/gin-gonic/gin"
)

type RoutesCollection struct{}

var CC = controllers.ControllerCollection{}

func RoutesGroup() *gin.Engine {
	r := gin.Default()

	api_v1 := r.Group("/api/v1")
	routesCollection := RoutesCollection{}

	routesCollection.AdminRoutes(api_v1)
	routesCollection.AdminRoutes(api_v1)
	routesCollection.UserRoutes(api_v1)
	routesCollection.ProductRoutes(api_v1)
	routesCollection.TransactionRoutes(api_v1)

	return r
}

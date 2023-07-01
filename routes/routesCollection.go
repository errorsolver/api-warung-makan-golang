package routes

import (
	"golang-wm-api/controllers"

	"github.com/gin-gonic/gin"
)

func RouterCollection() *gin.Engine {
	cC := controllers.ControllerCollection{}

	r := gin.Default()

	r.GET("/products", cC.GetProducts)
	r.GET("/product/:id", cC.GetProductById)
	r.POST("/product", cC.CreateProduct)
	r.PUT("/product/:id", cC.UpdateProduct)
	r.DELETE("/product", cC.DeleteProduct)

	r.GET("/user", cC.GetUsers)
	r.POST("/user", cC.CreateUser)

	r.GET("/transaction", cC.GetTransactionDetails)
	r.POST("/transaction", cC.CreateTransactionDetails)

	return r
}

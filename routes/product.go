package routes

import (
	"golang-wm-api/middleware"

	"github.com/gin-gonic/gin"
)

// type ProductRoutes struct{}

func (RoutesCollection) ProductRoutes(r *gin.RouterGroup) {
	r.GET("/products", CC.GetProducts)
	r.GET("/product/:id", CC.GetProductById)
	r.POST("/product", middleware.AdminAuth(), CC.CreateProduct)
	r.PUT("/product/:id", middleware.AdminAuth(), CC.UpdateProduct)
	r.DELETE("/product", middleware.AdminAuth(), CC.DeleteProduct)
}

// func ProductRouter() *gin.Engine {
// 	cPC := controllers.ProductControllers{}

// 	r := gin.Default()

// 	r.GET("/products", cPC.GetProducts)
// 	r.GET("/product/:id", cPC.GetProductById)
// 	r.POST("/product", cPC.CreateProduct)
// 	r.PUT("/product/:id", cPC.UpdateProduct)
// 	r.DELETE("/product", cPC.DeleteProduct)

// 	return r
// }

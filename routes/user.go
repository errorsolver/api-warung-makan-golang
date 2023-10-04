package routes

import (
	"golang-wm-api/middleware"

	"github.com/gin-gonic/gin"
)

// type UserRoutes struct{}

func (RoutesCollection) UserRoutes(r *gin.RouterGroup) {
	r.GET("/user", middleware.AdminAuth(), CC.GetUsers)
	r.GET("/user/:id", middleware.AdminAuth(), CC.GetUsersById)
	r.POST("/user", CC.CreateUser)
	r.PUT("/user/:id", CC.UpdateUser)
	r.DELETE("user/:id", middleware.AdminAuth(), CC.DeleteUser)
}

// import (
// 	"golang-wm-api/controllers"
// 	"github.com/gin-gonic/gin"
// )

// func UsersRouter() *gin.Engine {
// 	cPC := controllers.UserControllers{}
// 	r := gin.Default()

// 	r.GET("/", cPC.GetUsers)
// 	r.POST("/", cPC.CreateUser)

// 	return r
// }

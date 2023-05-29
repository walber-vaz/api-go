package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/api-go/src/controllers"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controllers.FindUserByID)
	r.GET("/user/email/:userEmail", controllers.FindUserByEmail)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:userId", controllers.UpdateUser)
	r.DELETE("/user/:userId", controllers.DeleteUser)
}

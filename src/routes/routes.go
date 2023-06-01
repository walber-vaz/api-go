package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/api-go/src/controllers"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to API Go",
		})
	})
	r.GET("/jobs", controllers.FindAllJobs)
	r.GET("/job/:id", controllers.FindJobId)
	r.POST("/job", controllers.CreateJob)
	r.PUT("/job/:id", controllers.UpdateJob)
	r.DELETE("/job/:id", controllers.DeleteJob)
}

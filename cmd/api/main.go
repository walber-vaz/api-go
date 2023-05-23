package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler() *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	return app
}

func main() {
	app := mainHandler()

	app.Run(":9000")
}

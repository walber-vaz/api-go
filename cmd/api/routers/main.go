package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/api-go/cmd/api/handlers"
)

func RouterHandler() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/")
	{
		v1.GET("jobs", handlers.Get)
	}

	return r
}

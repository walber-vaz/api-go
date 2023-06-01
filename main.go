package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/walber-vaz/api-go/src/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file" + err.Error())
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Error running server" + err.Error())
	}
}

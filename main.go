package main

import (
	"api/database"
	"api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	database.InitConnectionPool()

	app := gin.New()

	routes.ApiRoutes(app)

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		httpPort = "8080"
	}

	app.Run(":" + httpPort)
	fmt.Println("Server listening")
}

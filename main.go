package main

import (
	"nbapp/config"
	"nbapp/routes"

	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// config.InsertTeams()

	r := gin.Default()

	// Habilita requisições vindas do frontend no localhost:8081
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r)

	r.Run(":8080")
}

package routes

import (
	"nbapp/controllers"
	"nbapp/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rota do login
	router.POST("/login", controllers.Login)

	// Rotas proteginas
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Rotas de usuários (CRUD)
		userRoutes := router.Group("/users")
		{
			userRoutes.POST("/", controllers.CreateUser)
			userRoutes.GET("/", controllers.GetUsers)
			userRoutes.GET("/:id", controllers.GetUserByID)
			userRoutes.PUT("/:id", controllers.UpdateUser)
			userRoutes.DELETE("/:id", controllers.DeleteUser)
		}
	}
}

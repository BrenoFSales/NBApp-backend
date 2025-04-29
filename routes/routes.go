package routes

import (
	"nbapp/controllers"
	"nbapp/controllers/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rota do login
	router.POST("/login", controllers.Login)

	// Rota desprotegida temporariamente
	router.GET("/team-logo/:id", controllers.GetLogo)

	// Rotas proteginas
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Rotas de usuários (CRUD)
		userRoutes := protected.Group("/users")
		{
			userRoutes.POST("/", controllers.CreateUser)
			userRoutes.GET("/", controllers.GetUsers)
			userRoutes.GET("/:id", controllers.GetUserByID)
			userRoutes.PUT("/:id", controllers.UpdateUser)
			userRoutes.DELETE("/:id", controllers.DeleteUser)
		}

		teamRoutes := protected.Group("/teams")
		{
			teamRoutes.GET("/", controllers.GetTeams)
			teamRoutes.GET("/:id", controllers.GetTeamsById)
			teamRoutes.PUT("/upload-logo/:id", controllers.UploadLogoTeam)
			// teamRoutes.GET("/logo/:id", controllers.GetLogo)
		}
	}
}

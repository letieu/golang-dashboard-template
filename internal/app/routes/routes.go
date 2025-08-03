package routes

import (
	"chatpilot/app/internal/app/handlers"
	"chatpilot/app/internal/app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/static", "web/static")

	r.GET("/login", handlers.Login)
	r.GET("/logout", handlers.Logout)
	r.POST("/login-google-callback", handlers.LoginGoogleCallback)

	authGroup := r.Group("/")
	authGroup.Use(middlewares.AuthMiddleware())
	{
		authGroup.GET("/", handlers.Index)
		authGroup.POST("/agent", handlers.CreateNewAgent)
		authGroup.GET("/parts/agent/overview/:agentId", handlers.GetOverviewPart)
		authGroup.POST("/parts/agent/overview/:agentId", handlers.GetOverviewPart) // TODO: handle update detail
	}
}

package routes

import (
	"chatpilot/app/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/static", "web/static")
	r.GET("/", handlers.Index)
	r.GET("/prompt", handlers.Prompt)
	r.GET("/faqs", handlers.Faqs)
}

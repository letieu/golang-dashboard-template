package handlers

import (
	"chatpilot/app/internal/app/middlewares"
	"chatpilot/app/web/templates/layouts"
	"chatpilot/app/web/templates/pages"

	"github.com/gin-gonic/gin"
)

func Prompt(c *gin.Context) {
	c.Status(200)
	user := middlewares.CurrentUser{
		Username: "tieu",
		Email:    "uem",
	}
	layouts.DashboardLayout("Prompt", user, pages.PromptPage()).Render(c.Request.Context(), c.Writer)
}

package utils

import (
	"chatpilot/app/internal/app/middlewares"
	"chatpilot/app/web/templates/layouts"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderDashboardPage(content templ.Component, title string, c *gin.Context) {
	user, exist := c.Get("currentUser")
	if exist == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	}

	layouts.DashboardLayout(title, user.(middlewares.CurrentUser), content).Render(c.Request.Context(), c.Writer)
}

package utils

import (
	"chatpilot/app/web/templates/layouts"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func RenderDashboardPage(content templ.Component, title string, c *gin.Context) {
	user := GetCurrentUser(c)
	layouts.DashboardLayout(title, user, content).Render(c.Request.Context(), c.Writer)
}

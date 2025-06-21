package handlers

import (
	"chatpilot/app/internal/app/utils"
	"chatpilot/app/web/templates/pages"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	utils.RenderDashboardPage(pages.IndexPage(), "Home page", c)
}

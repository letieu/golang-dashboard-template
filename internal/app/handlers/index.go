package handlers

import (
	"chatpilot/app/web/templates/pages"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.Status(200)
	pages.IndexPage().Render(c.Request.Context(), c.Writer)
}

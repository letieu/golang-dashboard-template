package handlers

import (
	"chatpilot/app/web/templates/pages"

	"github.com/gin-gonic/gin"
)

func Faqs(c *gin.Context) {
	c.Status(200)
	pages.FaqsPage().Render(c.Request.Context(), c.Writer)
}

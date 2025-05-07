package handlers

import (
	"chatpilot/app/web/templates/pages"

	"github.com/gin-gonic/gin"
)

func Prompt(c *gin.Context) {
	c.Status(200)
	pages.PromptPage().Render(c.Request.Context(), c.Writer)
}

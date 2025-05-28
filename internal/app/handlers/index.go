package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/web/templates/pages"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db.DbPool.Exec("SELECT * from *")
	c.Status(200)
	pages.IndexPage().Render(c.Request.Context(), c.Writer)
}

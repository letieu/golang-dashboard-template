package user

import (
	"chatpilot/app/views/pages"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/", IndexPage)
}

func IndexPage(c *gin.Context) {
	c.Status(200)
	// views.Index().Render(c.Request.Context(), c.Writer)
	pages.HomePage().Render(c.Request.Context(), c.Writer)
}

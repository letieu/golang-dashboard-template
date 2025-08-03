package utils

import (
	"chatpilot/app/internal/app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) middlewares.CurrentUser {
	user, exist := c.Get("currentUser")
	if exist == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	}

	return user.(middlewares.CurrentUser)
}

package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/internal/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewAgent(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)
	_, err := db.CreateAgent(db.CreateAgentInput{
		Name:        "Agent 1",
		UserId:      currentUser.Id,
		Industry:    "",
		Personality: "",
	}, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/internal/app/utils"
	"chatpilot/app/web/templates/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)
	userAgents, err := db.GetAgentsByUserId(currentUser.Id, c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(userAgents) == 0 {
		utils.RenderDashboardPage(pages.NoAgentPage(currentUser), "Agent overview", c)
		return
	}

	var selectedIndex uint = 0
	utils.RenderDashboardPage(pages.AgentPage(
		selectedIndex,
		userAgents,
		currentUser,
	), "Agent overview", c)
}

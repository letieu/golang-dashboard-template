package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/internal/app/utils"
	"chatpilot/app/web/templates/pages"
	"log"
	"net/http"
	"strconv"

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
	var selectedId = userAgents[len(userAgents)-1].Id
	queryId, ok := c.GetQuery("agent")
	if ok == true {
		selectedId, err = strconv.ParseInt(queryId, 10, 64)
		if err != nil {
			log.Printf("%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	utils.RenderDashboardPage(pages.AgentPage(
		selectedId,
		userAgents,
		currentUser,
	), "Agent overview", c)
}

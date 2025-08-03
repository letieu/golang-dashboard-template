package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/internal/app/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"chatpilot/app/web/templates/parts"
	"github.com/gin-gonic/gin"
)

func CreateNewAgent(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)
	newId, err := db.CreateAgent(db.CreateAgentInput{
		Name:        generateMilitaryAgent(),
		UserId:      currentUser.Id,
		Industry:    "",
		Personality: "",
	}, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/?agent="+strconv.Itoa(int(newId)))
}

// /parts/agent/overview/123
func GetOverviewPart(c *gin.Context) {
	currentUser := utils.GetCurrentUser(c)
	agentIdStr, ok := c.Params.Get("agentId")
	if ok == false {
		parts.ErrorInfo("Do not have agentId").Render(c.Request.Context(), c.Writer)
		return
	}
	agentId, err := strconv.Atoi(agentIdStr)
	if err != nil {
		parts.ErrorInfo("Invalid agentId").Render(c.Request.Context(), c.Writer)
		return
	}

	agent, err := db.GetAgentById(agentId, int(currentUser.Id), c.Request.Context())
	if err != nil {
		parts.ErrorInfo(err.Error()).Render(c.Request.Context(), c.Writer)
		return
	}

	parts.AgentOverview(agent).Render(c, c.Writer)
}

func generateMilitaryAgent() string {
	prefixes := []string{
		"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel",
		"India", "Juliet", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa",
		"Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu",
	}

	prefix := prefixes[rand.Intn(len(prefixes))]
	number := rand.Intn(999) + 1

	return fmt.Sprintf("%s-%03d", prefix, number)
}

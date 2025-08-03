package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/internal/app/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

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

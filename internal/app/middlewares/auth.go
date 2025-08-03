package middlewares

import (
	"chatpilot/app/internal/app/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrentUser struct {
	Id       int64
	Username string
	Email    string
}

func validateToken(sessionToken string) (CurrentUser, error) {
	var user CurrentUser

	// TODO: handle expire
	query := `
        SELECT u.id, u.username, u.email
        FROM session s
        JOIN user u ON s.user_id = u.id
        WHERE s.token = ?
    `
	row := db.DbPool.QueryRow(query, sessionToken)
	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session_token")
		if err != nil || token == "" {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}

		user, err := validateToken(token)
		if err != nil {
			fmt.Printf("Error %v \n", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

package middlewares

import (
	"chatpilot/app/internal/app/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrentUser struct {
	Username string
	Email    string
	Id       uint
}

func validateToken(sessionToken string) (CurrentUser, error) {
	var user CurrentUser

	// TODO: handle expire
	query := `
        SELECT u.rowid, u.username, u.email
        FROM session s
        JOIN user u ON s.user_email = u.email
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
			return
		}

		user, err := validateToken(token)
		if err != nil {
			fmt.Printf("Error %v \n", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

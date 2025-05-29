package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/web/templates/pages"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

func Login(c *gin.Context) {
	c.Status(200)
	pages.LoginPage().Render(c.Request.Context(), c.Writer)
}

func LoginGoogleCallback(c *gin.Context) {
	defer c.Request.Body.Close()

	// verify cookie
	csrfTokenCookie, err := c.Cookie("g_csrf_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get csrf cookie"})
		return
	}

	csrfTokenBody := c.PostForm("g_csrf_token")
	if csrfTokenBody != csrfTokenCookie {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to verify double submit cookie."})
		return
	}

	// verify token
	validator, err := idtoken.NewValidator(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error to create idtoken validator"})
		return
	}

	credential := c.PostForm("credential")
	payload, err := validator.Validate(c, credential, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error validate idtoken"})
		return
	}

	// upsert to db
	username := payload.Claims["name"]
	email := payload.Claims["email"]

	_, err = db.DbPool.ExecContext(c, `
        INSERT INTO user (email, username)
        VALUES (?, ?)
        ON CONFLICT(email) DO UPDATE SET
            username = excluded.username
    `, email, username)

	if err != nil {
		fmt.Println("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// gen session
	sessionToken, err := generateSessionToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when create new session token"})
		return
	}
	_, err = db.DbPool.Exec(`
		INSERT INTO sessions (token, user_email) VALUES (?, ?)
	`, sessionToken, email)
	if err != nil {
		fmt.Println("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.SetCookie("session_token", sessionToken, 3600*24*7, "/", "", false, true)

	// set cookie
	c.Redirect(301, "/")
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 32) // 32 bytes = 64 hex characters
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

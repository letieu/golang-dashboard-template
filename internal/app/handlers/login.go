package handlers

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/web/templates/pages"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const sessionDuration = time.Hour * 24 * 30 // 30 days

func Login(c *gin.Context) {
	pages.LoginPage().Render(c.Request.Context(), c.Writer)
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil || token == "" {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	db.DbPool.ExecContext(c, `
		DELETE FROM session WHERE token = ? 
	`, token)
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGoogleCallback(c *gin.Context) {
	defer c.Request.Body.Close()

	if err := verifyCSRF(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, err := validateIDToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, email, err := extractUserInfo(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = findOrCreateUser(email, username, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := initSession(email, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func generateSessionToken() (string, error) {
	bytes := make([]byte, 32) // 32 bytes = 64 hex characters
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func generateAgent(userId uint, ctx context.Context) error {
	tenantIds, err := db.GetUserTenantIds(userId, ctx)
	if err != nil {
		return err
	}

	var tenantId uint
	if len(tenantIds) == 0 {
		tenantId, err = db.CreateDefaultTenantForUser(userId, ctx)
		if err != nil {
			return err
		}
	} else {
		tenantId = tenantIds[0]
	}

	agentIds, err := db.GetAgentIdsByTenantId(tenantId, ctx)
	if err != nil {
		return err
	}
	if len(agentIds) == 0 {
		if err := db.CreateDefaultAgentForTenant(tenantId, ctx); err != nil {
			return err
		}
	}

	return nil
}

func verifyCSRF(c *gin.Context) error {
	csrfCookie, err := c.Cookie("g_csrf_token")
	if err != nil {
		return fmt.Errorf("failed to get CSRF cookie")
	}
	csrfBody := c.PostForm("g_csrf_token")
	if csrfBody != csrfCookie {
		return fmt.Errorf("failed to verify double submit cookie")
	}
	return nil
}

func validateIDToken(c *gin.Context) (*idtoken.Payload, error) {
	validator, err := idtoken.NewValidator(c)
	if err != nil {
		return nil, fmt.Errorf("failed to create ID token validator")
	}
	credential := c.PostForm("credential")
	payload, err := validator.Validate(c, credential, "")
	if err != nil {
		return nil, fmt.Errorf("failed to validate ID token")
	}
	return payload, nil
}

func extractUserInfo(payload *idtoken.Payload) (username, email string, err error) {
	rawUsername, ok1 := payload.Claims["name"].(string)
	rawEmail, ok2 := payload.Claims["email"].(string)

	if !ok1 || !ok2 {
		return "", "", fmt.Errorf("failed to extract username or email")
	}
	return rawUsername, rawEmail, nil
}

func findOrCreateUser(email, username string, c *gin.Context) (uint, error) {
	userId, err := db.GetUserIdByEmail(email, c)
	if err == nil {
		return userId, nil
	}

	// Create new user
	if err := db.UpsertUser(email, username, c); err != nil {
		return 0, fmt.Errorf("failed to create user")
	}

	userId, err = db.GetUserIdByEmail(email, c)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve user after creation")
	}

	// Create tenant + agent
	if err := generateAgent(userId, c); err != nil {
		return 0, fmt.Errorf("failed to create default agent for user")
	}

	return userId, nil
}

func initSession(email string, c *gin.Context) error {
	sessionToken, err := generateSessionToken()
	if err != nil {
		return fmt.Errorf("failed to generate session token")
	}

	expireAt := time.Now().Add(sessionDuration)
	if err := db.SetSession(email, sessionToken, expireAt); err != nil {
		return fmt.Errorf("failed to store session: %v", err)
	}

	maxAge := int(time.Until(expireAt).Seconds())
	c.SetCookie("session_token", sessionToken, maxAge, "/", "", false, true)
	return nil
}

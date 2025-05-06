package main

import (
	"chatpilot/app/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("Starting server...")

	user.RegisterRoutes(r.Group("/users"))

	// static files
	r.Static("/static", "./static")

	r.Run(":8080")
}

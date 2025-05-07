package main

import (
	"chatpilot/app/internal/app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}

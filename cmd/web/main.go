package main

import (
	"chatpilot/app/internal/app/db"
	"chatpilot/app/internal/app/routes"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Starting server...")

	var err error
	db.DbPool, err = sql.Open("sqlite3", "saleAI.db")
	if err != nil {
		return
	}
	defer db.DbPool.Close()

	if err := db.DbPool.Ping(); err != nil {
		log.Fatalf("Can't ping db")
	}

	fmt.Println("SQLite DB connected successfully.")

	fmt.Println("Setting routes...")

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8000")
}

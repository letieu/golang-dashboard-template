package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pool, err := sql.Open("sqlite3", "saleAI.db")
	if err != nil {
		log.Fatalf("Failed to open SQLite DB: %v", err)
		return
	}

	defer pool.Close()

	if err := pool.Ping(); err != nil {
		log.Fatalf("Can't ping db")
	}

	fmt.Println("SQLite DB connected successfully.")

	query := loadInitQuery()
	_, err = pool.Exec(query)
	if err != nil {
		log.Fatalf("Error exec init query: %v", err)
	}

	fmt.Println("Init data successfully")
}

func loadInitQuery() string {
	file, err := os.ReadFile("init_db.sql")
	if err != nil {
		log.Fatalf("Error to read file %v", err)
	}

	return string(file)
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
)

func main() {
	// Database connection parameters
	db_host := "192.168.3.21"
	db_port := "1521"
	db_service := "pdb1"
	db_user := "testuser"
	db_password := "Test1234"

	// Create connection string
	connStr := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", db_user, db_password, db_host, db_port, db_service)

	// Open database connection
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Successfully connected to Oracle Database 21c!")

	// Optional: Query some basic information
	var version string
	err = db.QueryRow("SELECT version FROM v$instance").Scan(&version)
	if err != nil {
		log.Println("Could not query database version:", err)
	} else {
		fmt.Println("Database version:", version)
	}
}

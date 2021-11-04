package main

import (
	"fmt"
	"log"
	"os"

	pg "github.com/infamousjoeg/qwik-e-mart/pkg/postgres"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = 5432
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func main() {
	// Connect to the database
	db, err := pg.Connect(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Query the database
	rows, err := pg.Query(db)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Print connection information received in env vars
	fmt.Println("-----------------------------------")
	fmt.Printf("Connected successfully to %s\n", host)
	fmt.Printf("Database Username: %s\n", user)
	fmt.Printf("Database Password: %s\n", password)
	fmt.Println("-----------------------------------")
	fmt.Println("")

	// Print all rows returned
	err = pg.PrintRows(rows)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Close query and connection
	pg.CloseQuery(rows)
	pg.Close(db)
}

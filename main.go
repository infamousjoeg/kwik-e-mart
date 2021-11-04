package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = 5432
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
	emptyEnv []string
)

func checkEnvValues() error {
	if host == "" {
		emptyEnv = append(emptyEnv, "DB_HOST")
	}
	if user == "" {
		emptyEnv = append(emptyEnv, "DB_USERNAME")
	}
	if password == "" {
		emptyEnv = append(emptyEnv, "DB_PASSWORD")
	}
	if dbname == "" {
		emptyEnv = append(emptyEnv, "DB_NAME")
	}
	if emptyEnv == nil {
		return nil
	}

	return fmt.Errorf("required environment variable(s) missing: %s", strings.Join(emptyEnv, ","))
}

func main() {
	// Check required environment variable values
	err := checkEnvValues()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Database connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	// Open database
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatalf("an error occured while connecting to database: %s", err)
	}

	// Close database
	defer db.Close()

	// Query database for returned rows
	rows, err := db.Query(`SELECT "id", "first_name", "last_name", "pmt_type" FROM "customers"`)
	if err != nil {
		log.Fatalf("error while querying database: %s", err)
	}

	// Close query connection
	defer rows.Close()

	fmt.Println("-----------------------------------")
	fmt.Printf("Connected successfully to %s\n", host)
	fmt.Printf("Database Username: %s\n", user)
	fmt.Printf("Database Password: %s\n", password)
	fmt.Println("-----------------------------------")
	fmt.Println("")
	fmt.Println("id, first_name, last_name, pmt_type")
	fmt.Println("-----------------------------------")

	// Print returned rows
	for rows.Next() {
		var id int
		var first_name string
		var last_name string
		var pmt_type string

		err = rows.Scan(&id, &first_name, &last_name, &pmt_type)
		if err != nil {
			log.Fatalf("error scanning returned row data: %s", err)
		}

		fmt.Printf("%d, %s, %s, %s\n", id, first_name, last_name, pmt_type)
	}
}

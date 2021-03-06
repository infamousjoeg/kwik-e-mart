package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	pg "github.com/infamousjoeg/kwik-e-mart/pkg/postgres"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = 5432
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

// Index is the default route
func Index(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintln(w, "-------------------------------------------------------")
	fmt.Fprintf(w, "Connected successfully to %s\n", host)
	fmt.Fprintf(w, "Database Username: %s\n", user)
	fmt.Fprintf(w, "Database Password: %s\n", password)
	fmt.Fprintln(w, "-------------------------------------------------------")
	fmt.Fprintln(w, "")

	// Print all rows returned
	err = pg.PrintRows(w, rows)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Close query and connection
	pg.CloseQuery(rows)
	pg.Close(db)
}

func main() {
	// Create new gorilla/mux router
	router := mux.NewRouter()

	// Add routes
	router.HandleFunc("/", Index).Methods("GET")

	// Start server
	fmt.Println("-----------------------------------")
	fmt.Println("Starting server on port 8080")
	fmt.Println("-----------------------------------")
	fmt.Println("Browse to: http://<dns>:8080")
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

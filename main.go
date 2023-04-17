package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	conjur "github.com/aharriscybr/kwik-e-mart/pkg/conjur"
	pg "github.com/aharriscybr/kwik-e-mart/pkg/postgres"
	"github.com/gorilla/mux"
)

var (
	// Authn Data
	token   = os.Getenv("CONJUR_TOKEN_PATH")
	baseUri = os.Getenv("CONJUR_BASE")
	accnt   = os.Getenv("CONJUR_ACCOUNT")
	safe    = os.Getenv("CONJUR_SAFE")
	query   = os.Getenv("CONJUR_QUERY")

)

// Index is the default route
func Index(w http.ResponseWriter, r *http.Request) {


	secrets := conjur.GetData(baseUri, token, accnt, safe, query)

	/*
	*	secrets[0] is the user
	*	secrets[1] is the pass
	*	secrets[2] is the port
	*	secrets[3] is the dbname
	*	secrets[4] is the address
	*/

	portCon := 0
	
	_, err := fmt.Sscan(secrets[2], portCon)
	if err != nil {
		log.Println(err)
	}

	// Connect to the database
	db, err := pg.Connect(secrets[4], portCon, secrets[0], secrets[1], secrets[3])
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
	fmt.Fprintf(w, "Connected successfully to %s\n", secrets[4])
	fmt.Fprintf(w, "Database Username: %s\n", secrets[0])
	fmt.Fprintf(w, "Database Password: %s\n", secrets[1])
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

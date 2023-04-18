package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var (
	emptyEnv []string
)

func checkEnvValues(host string, user string, password string, dbname string) error {
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

// Connect completes a connection to a PostgreSQL database host
func Connect(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	// Check required environment variable values
	err := checkEnvValues(host, user, password, dbname)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	// Database connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s", host, port, user, password)

	// Open database
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, fmt.Errorf("an error occured while connecting to database: %s", err)
	}

	return db, nil
}

// Close the database connection using by calling this function
func Close(db *sql.DB) {
	// Close database
	defer db.Close()
}

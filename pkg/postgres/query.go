package postgres

import (
	"database/sql"
	"fmt"
)

// Query runs a transactional statement on a PostgreSQL database
func Query(db *sql.DB) (*sql.Rows, error) {
	// Query database for returned rows
	rows, err := db.Query(`SELECT "id", "first_name", "last_name", "pmt_type" FROM "customers"`)
	if err != nil {
		return nil, fmt.Errorf("error while querying database: %s", err)
	}

	return rows, nil
}

// CloseQuery closes the query connection
func CloseQuery(rows *sql.Rows) {
	// Close query connection
	defer rows.Close()
}

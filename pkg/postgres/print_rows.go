package postgres

import (
	"database/sql"
	"fmt"
)

// PrintRows prints out the returned rows provided to the function
func PrintRows(rows *sql.Rows) error {
	fmt.Println("id, first_name, last_name, pmt_type")
	fmt.Println("-----------------------------------")

	// Print returned rows
	for rows.Next() {
		var id int
		var first_name string
		var last_name string
		var pmt_type string

		err := rows.Scan(&id, &first_name, &last_name, &pmt_type)
		if err != nil {
			return fmt.Errorf("error scanning returned row data: %s", err)
		}

		fmt.Printf("%d, %s, %s, %s\n", id, first_name, last_name, pmt_type)
	}

	return nil
}

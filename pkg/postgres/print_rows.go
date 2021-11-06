package postgres

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/tabwriter"
)

// PrintRows prints out the returned rows provided to the function
func PrintRows(w http.ResponseWriter, rows *sql.Rows) error {
	t := tabwriter.NewWriter(w, 15, 4, 1, ' ', 0)
	fmt.Fprintln(t, "id\tfirst_name\tlast_name\tpmt_type")
	fmt.Fprintln(t, "-------------------------------------------------------")
	t.Flush()

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
		t := tabwriter.NewWriter(w, 15, 4, 1, ' ', 0)
		fmt.Fprintf(t, "%d\t%s\t%s\t%s\n", id, first_name, last_name, pmt_type)
		t.Flush()
	}

	return nil
}

package ticket

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

const schema = `
CREATE TABLE IF NOT EXISTS ticket (
	id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY
);
`

// We don't close the connection here as it's intended to be shared.
func LoadSchema(conn *sqlx.DB) {
	_, err := conn.Exec(schema)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading 'ticket' schema: %v", err)
	}
}

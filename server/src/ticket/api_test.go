package ticket

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

func cleanup(conn *sqlx.DB, id int64) {
	conn.MustExec("DELETE FROM ticket WHERE id = $1;", id)
}

func connect() *sqlx.DB {
	conn_string := fmt.Sprintf("postgres://pos:%s@database-service/pos", os.Getenv("USER_PASSWORD"))
	return sqlx.MustConnect("pgx", conn_string)
}

func TestTicketIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test.")
	}

	conn := connect()
	defer conn.Close()

	t.Run("Insert", func(t *testing.T) {
		t_1 := Ticket{OrderNumber: 1, TableNumber: 1, Server: "TestServer"}

		id := Insert(conn, t_1)

		defer cleanup(conn, id)
	})
}

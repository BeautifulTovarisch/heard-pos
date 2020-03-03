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

func fetch_ticket(conn *sqlx.DB, id int64) Ticket {
	var result Ticket
	query := "SELECT server, order_no, table_no, opened, closed FROM ticket WHERE id = $1;"
	conn.QueryRowx(query, id).StructScan(&result)

	return result
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

	t.Run("Order Count", func(t *testing.T) {
		t_1 := Ticket{Server: "TestServer", TableNumber: 1}
		t_2 := Ticket{Server: "TestServer2", TableNumber: 2}

		id_1 := Insert(conn, t_1)
		id_2 := Insert(conn, t_2)

		t_2 = fetch_ticket(conn, id_2)

		if t_2.OrderNumber < 2 {
			t.Errorf("Expected OrderNumber >= 2. Got: %d", t_2.OrderNumber)
		}

		defer cleanup(conn, id_2)
		defer cleanup(conn, id_1)
	})
}

// api.go /api/ticket
//
// RESTful endpoints to abstract database interaction for Tickets.
//
package ticket

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func Routes(conn *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	return router
}

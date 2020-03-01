package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"

	//	"heard/menu"
	"heard/ticket"
)

func Routes(conn *sqlx.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
	)

	router.Route("/v0", func(r chi.Router) {
		//	r.Mount("/api/menu", menu.Routes())
		r.Mount("/api/ticket", ticket.Routes(conn))
	})

	return router
}

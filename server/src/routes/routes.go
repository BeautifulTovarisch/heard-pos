package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	//	"heard/menu"
	"heard/ticket"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
	)

	router.Route("/v0", func(r chi.Router) {
		//	r.Mount("/api/menu", menu.Routes())
		r.Mount("/api/ticket", ticket.Routes())
	})

	return router
}

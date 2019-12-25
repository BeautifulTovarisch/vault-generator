package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"vault-generator/vault"
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
		r.Mount("/api/vault", vault.Routes())
	})

	return router
}

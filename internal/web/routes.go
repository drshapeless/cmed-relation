package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", app.home)

	r.Route("/search", func(r chi.Router) {
		r.Post("/disease", app.searchDisease)
		r.Post("/herb", app.searchHerb)
		r.Post("/formula", app.searchFormula)
	})

	return r
}

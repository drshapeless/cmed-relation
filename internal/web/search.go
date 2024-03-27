package web

import (
	"cmed-relation/internal/data"
	"cmed-relation/internal/templates"
	"net/http"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	err := templates.Base("home", templates.SearchPage()).Render(r.Context(), w)
	if err != nil {
		app.serverError(err, w)
		return
	}
}

func (app *Application) searchDisease(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(err, w)
		return
	}

	name := r.FormValue("name")

	sources, err := data.GetSourceByDisease(name, app.DB)
	if err != nil {
		app.serverError(err, w)
		return
	}

	err = templates.SourceResult(sources).Render(r.Context(), w)
	if err != nil {
		app.serverError(err, w)
		return
	}
}

func (app *Application) searchHerb(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(err, w)
		return
	}

	name := r.FormValue("name")

	sources, err := data.GetSourceByHerb(name, app.DB)
	if err != nil {
		app.serverError(err, w)
		return
	}

	err = templates.SourceResult(sources).Render(r.Context(), w)
	if err != nil {
		app.serverError(err, w)
		return
	}
}

func (app *Application) searchFormula(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(err, w)
		return
	}

	name := r.FormValue("name")

	sources, err := data.GetSourceByFormula(name, app.DB)
	if err != nil {
		app.serverError(err, w)
		return
	}

	err = templates.SourceResult(sources).Render(r.Context(), w)
	if err != nil {
		app.serverError(err, w)
		return
	}
}

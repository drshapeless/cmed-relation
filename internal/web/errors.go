package web

import (
	"cmed-relation/internal/logger"
	"net/http"
)

func (app *Application) serverError(err error, w http.ResponseWriter) {
	app.Logger.Output(3, logger.LogLevelError, err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

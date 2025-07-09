package main

import "net/http"

func (app *application) router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /confess", app.confess)
	mux.HandleFunc("POST /confess", app.confessPost)
	mux.HandleFunc("GET /confession/{id}", app.confessionViewId)
	mux.HandleFunc("GET /confession/all", app.confessionViewAll)

	return mux
}

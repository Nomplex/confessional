package main

import "net/http"

func (app *application) router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /confession", app.confession)
	mux.HandleFunc("POST /confession", app.confessionPost)
	mux.HandleFunc("GET /confession/{id}", app.confessionViewId)
	mux.HandleFunc("GET /confession/all", app.confessionViewAll)

	return mux
}

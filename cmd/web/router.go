package main

import "net/http"

func (app *application) router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)

	return mux
}

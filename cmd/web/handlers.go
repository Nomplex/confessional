package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	confessions, err := app.confessions.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := &templateData{confessions}
	app.render(w, r, http.StatusOK, "home.tmpl.html", *data)
}

func (app *application) confess(w http.ResponseWriter, r *http.Request) {
	data := &templateData{}
	app.render(w, r, http.StatusOK, "confess.tmpl.html", *data)
}

func (app *application) confessPost(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("New Post Submission Success")
}

func (app *application) confessionViewId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "View Confession %s", id)
}

func (app *application) confessionViewAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Viewing All Confessions")
}

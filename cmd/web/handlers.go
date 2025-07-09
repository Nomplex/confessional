package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func (app *application) confess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "New Confession")
}

func (app *application) confessPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post New Confession")
}

func (app *application) confessionViewId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "View Confession %s", id)
}

func (app *application) confessionViewAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Viewing All Confessions")
}

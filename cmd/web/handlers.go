package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	confessions, err := app.confessions.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{}
	data.Confessions = confessions

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) confess(w http.ResponseWriter, r *http.Request) {
	data := templateData{}
	app.render(w, r, http.StatusOK, "confess.tmpl.html", data)
}

type confessPostForm struct {
	Title       string
	Content     string
	FieldErrors map[string]string
}

func (app *application) confessPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := confessPostForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(form.Title) > 25 {
		form.FieldErrors["title"] = "This field cannot be more than 25 characters"
	}

	if strings.TrimSpace(form.Content) == "" {
		form.FieldErrors["content"] = "This field cannot be blank"
	}

	if len(form.FieldErrors) > 0 {
		data := templateData{}
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "confess.tmpl.html", data)
		return
	}

	id, err := app.confessions.Insert(form.Title, form.Content)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/confession/%d", id), http.StatusSeeOther)
}

func (app *application) confessionViewId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "View Confession %s", id)
}

func (app *application) confessionViewAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Viewing All Confessions")
}

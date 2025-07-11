package main

import (
	"html/template"
	"path/filepath"
)

type templateData struct{}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Gives us a slice of all filepaths that match the pattern
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	// Iterate over the slice
	for _, page := range pages {
		// Gives us the file name
		name := filepath.Base(page)

		// In the future if I need template functions I can edit this
		ts, err := template.New(name).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Then we layer on any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Then we layer on the page
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Finally apply the template set to the cache
		cache[name] = ts
	}

	return cache, nil
}

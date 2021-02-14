package controllers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	)

func StartingPage(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	path := filepath.Join("public", "html", "page.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)

	}
}
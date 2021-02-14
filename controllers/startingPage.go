package controllers

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func StartingPage(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	path := filepath.Join("client", "main-page.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
	}
}

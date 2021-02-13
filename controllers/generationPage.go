package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"../models"
)

func GenerationPage(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	landscape, err := models.NewLandscape(10, 10)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = json.NewEncoder(rw).Encode(landscape)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
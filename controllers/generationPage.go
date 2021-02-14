package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"../models"
)

func GenerationPage(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	landscape, err := models.NewLandscape(20, 20)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	landscape.PerlinNoiseGeneration(50, 3, 0.5)

	err = json.NewEncoder(rw).Encode(landscape)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

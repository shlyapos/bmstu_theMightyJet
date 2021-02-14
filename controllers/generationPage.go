package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"

	"../models"
)

func GenerationPage(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	landscape, err := models.NewLandscape(32, 32)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	landscape.PerlinNoiseGeneration(50,4, 0.5)
	landscape.ToJPG("temp.png", 10)
	err = json.NewEncoder(rw).Encode(landscape)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

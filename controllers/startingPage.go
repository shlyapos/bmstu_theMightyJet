package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StartingPage(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	text := "Похуй нахуй похуй нахуй!"

	fmt.Fprint(rw, text)
}

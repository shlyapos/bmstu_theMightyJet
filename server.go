package main

import (
	"github.com/julienschmidt/httprouter"
	"flag"
	"net/http"
	"strconv"

	"./controllers"
)

type Server struct {
	port int
	router *httprouter.Router
}

func startServer(port int) {
	var server Server
	server.port = port
	server.router = httprouter.New()
	server.setRoutes()

	err := http.ListenAndServe("localhost:" + strconv.Itoa(server.port), server.router)
	if err != nil {
		panic(err)
	}
}

func (server *Server) setRoutes() {
	server.router.ServeFiles("/public/*filepath", http.Dir("public"))

	server.router.GET("/", controllers.StartingPage)
	server.router.GET("/generate", controllers.GenerationPage)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 4444, "Server port")
	flag.Parse()

	startServer(port)
}
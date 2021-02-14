package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"./controllers"
)

type Server struct {
	port   int
	router *httprouter.Router
}

func startServer(port int) {
	var server Server
	server.port = port
	server.router = httprouter.New()
	server.setRoutes()

	fmt.Println("Поехали, ковбой!")

	err := http.ListenAndServe("localhost:"+strconv.Itoa(server.port), server.router)
	if err != nil {
		panic(err)
	}
}

func (server *Server) setRoutes() {
	server.router.ServeFiles("/client/*filepath", http.Dir("client"))

	server.router.GET("/", controllers.StartingPage)
	server.router.GET("/generate", controllers.GenerationPage)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 4444, "Server port")
	flag.Parse()

	startServer(port)
}

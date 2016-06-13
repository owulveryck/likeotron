package server

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
	"net/http"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/json/{jsonContent}").
		Name("Dynamic").
		HandlerFunc(GetJson)

	router.
		Methods("GET").
		Path("/echo").
		Name("WebSocket").
		Handler(websocket.Handler(echoHandler))

	router.
		Methods("GET").
		PathPrefix("/").
		Name("Static").
		Handler(http.FileServer(http.Dir("./htdocs")))
	return router

}

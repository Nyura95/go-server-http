package server

import (
	"log"
	"messenger/config"
	"messenger/server/middleware"
	v1 "messenger/server/v1"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Start for start the http server
func Start() {
	config := config.GetConfig()
	r := mux.NewRouter()

	// Cors auth ajax
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	headersOk := handlers.AllowedHeaders([]string{"authorization"})

	// Middleware
	r.Use(middleware.BasicHeader)
	r.Use(middleware.Logger)

	// For request ajax
	r.HandleFunc("/", nil).Methods("OPTIONS")

	// Your endpoint
	r.HandleFunc("/api/v1/conversation", v1.GetConvo).Methods("POST")

	log.Println("Start http server on :" + strconv.Itoa(config.Port))
	http.ListenAndServe(":"+strconv.Itoa(config.Port), handlers.CORS(originsOk, headersOk, methodsOk)(r))
}

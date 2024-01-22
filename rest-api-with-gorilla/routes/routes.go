package routes

import (
	"github.com/gorilla/mux"
	"restapi/handlers"
)

func SetUp(router *mux.Router) {
	router.HandleFunc("/scooter", handlers.GetAll).Methods("GET")
	router.HandleFunc("/scooter/{id}", handlers.GetById).Methods("GET")
}

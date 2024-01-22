package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restapi/routes"
)

func main() {
	fmt.Println("API server starting..")
	router := mux.NewRouter()
	fmt.Println("Setting up routes..")
	routes.SetUp(router)
	fmt.Println("HTTP server started listening..")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

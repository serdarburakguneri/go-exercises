package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/models"
)

func fakeData() []models.Scooter {
	return []models.Scooter{
		{Id: uuid.MustParse("56ff36c1-b913-45f6-8fc3-d238569b6995"), Model: "Bird", SerialNumber: "ABC123"},
		{Id: uuid.MustParse("015aaec9-d8a1-4911-8271-2297b72e5ac6"), Model: "Bird", SerialNumber: "ABC124"},
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Received GET request")
	scooters := fakeData()
	json.NewEncoder(w).Encode(scooters)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := uuid.MustParse(params["id"])
	fmt.Printf("Received GET request for %s", id)

	//for the sake of simplicity, let's keep this searching logic in the handler
	scooters := fakeData()
	for _, scooter := range scooters {
		if scooter.Id == id {
			json.NewEncoder(w).Encode(scooter)
			return
		}
	}

	http.Error(w, fmt.Sprintf("A scooter with id:%s is not present", id), http.StatusNotFound)
}

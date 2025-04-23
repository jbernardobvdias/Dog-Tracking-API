package handler

import (
	"dog-tracking/models"
	"encoding/json"
	"net/http"
)

func GetAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs := models.Dogs

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dogs); err != nil {
		http.Error(w, "Failed to encode dogs data", http.StatusInternalServerError)
		return
	}
}

func GetDogById(w http.ResponseWriter, r *http.Request) {

}

func AddDog(w http.ResponseWriter, r *http.Request) {

}

func UpdateDog(w http.ResponseWriter, r *http.Request) {

}

func DeleteDog(w http.ResponseWriter, r *http.Request) {

}

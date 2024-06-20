package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receipts = map[string]Receipt{}

// Generates an ID for a receipt and returns the ID
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Generate ID
	id := uuid.New().String()

	// Save receipt to local map
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)

	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	receipts[id] = receipt

	// Send response
	response := ID{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Calculates and returns score of receipt given an ID
func CalculatePointsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if receipt, ok := receipts[params["id"]]; ok {
		score, _ := calculatePoints(receipt)

		response := Points{
			Points: score,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	http.Error(w, "No receipt found for that id", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", ProcessReceiptHandler).Methods(http.MethodPost)
	router.HandleFunc("/receipts/{id}/points", CalculatePointsHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

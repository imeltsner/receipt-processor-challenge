package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receipts []Receipt

// Generates an ID for a receipt and returns the ID
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Generate ID and send response
	id := uuid.New().String()

	response := ID{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Save receipt with new ID
	var receipt Receipt
	json.NewDecoder(r.Body).Decode(&receipt)
	receipt.ID = id
	receipts = append(receipts, receipt)
}

// Calculates and returns score of receipt given an ID
func CalculatePointsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, receipt := range receipts {
		if receipt.ID == params["id"] {
			score, _ := calculatePoints(receipt)

			response := Points{
				Points: score,
			}

			json.NewEncoder(w).Encode(response)
			return
		}

	}

	http.Error(w, "No receipt found for that id", http.StatusBadRequest)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", ProcessReceiptHandler).Methods(http.MethodPost)
	router.HandleFunc("/receipts/{id}/points", CalculatePointsHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receipts = map[string]Receipt{}

// Generates a deterministic ID for a receipt
func generateReceiptID(receipt Receipt) string {
	// Create unique string from receipt details and init namespace
	uniqueString := receipt.Retailer + receipt.PurchaseDate + receipt.Total
	namespace := uuid.UUID{}

	// Generate uuid based on unique string
	uuid := uuid.NewSHA1(namespace, []byte(uniqueString))

	return uuid.String()
}

// Generates an ID for a receipt and returns the ID
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Decode JSON
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Validate all fields present
	validate := validator.New()
	err = validate.Struct(receipt)
	if err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Generate ID and save new receipt to local memory
	id := generateReceiptID(receipt)
	if _, ok := receipts[id]; ok {
		http.Error(w, "The receipt has already been scanned", http.StatusBadRequest)
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
	router.HandleFunc("/receipts/{id}/points", CalculatePointsHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}

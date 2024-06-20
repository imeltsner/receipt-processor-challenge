package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Generates an ID for a receipt and returns the ID
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()

	response := ID{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", ProcessReceiptHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}

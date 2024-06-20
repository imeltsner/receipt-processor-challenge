package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

func deserialzeJSON(filepath string) (Receipt, error) {
	var receipt Receipt

	// Open JSON file
	file, err := os.Open(filepath)
	if err != nil {
		return receipt, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	// Read file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		return receipt, fmt.Errorf("unable to open file: %v", err)
	}

	// Deserialize struct
	err = json.Unmarshal(bytes, &receipt)
	if err != nil {
		return receipt, fmt.Errorf("unable to deserialize JSON: %v", err)
	}

	return receipt, nil
}

// Calculates the score associated with a given receipt
func calculateScore(receipt Receipt) (int, error) {
	score := 0

	// Adds all alphanumeric characters of retailer to score
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			score += 1
		}
	}

}

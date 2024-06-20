package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

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

func main() {
	receipt, err := deserialzeJSON("mocks/market-receipt.json")

	if err != nil {
		fmt.Errorf("error deserializing JSON: %v", err)
	}

	score, err := calculateScore(receipt)

	if err != nil {
		fmt.Errorf("unable to calculate score: %v", err)
	}

	fmt.Printf("Total score: %v", score)
}

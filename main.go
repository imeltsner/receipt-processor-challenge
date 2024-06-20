package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"` // TODO fix datatype
	PurchaseTime string `json:"purchaseTime"` // TODO fix datatype
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

// // Calculates the score associated with a given receipt
// func calculateScore(receipt Receipt) (int, error) {
// 	score := 0

// 	// Adds 1 point for each alphanumeric characters of retailer
// 	for _, c := range receipt.Retailer {
// 		if unicode.IsLetter(c) || unicode.IsDigit(c) {
// 			score += 1
// 		}
// 	}

// 	// Adds 50 points if total is round dollar amount
// 	if int(receipt.Total*100)%100 == 0 {
// 		score += 50
// 	}

// 	// Add 25 points if total is multiple of 0.25
// 	if int(receipt.Total*100)%25 == 0 {
// 		score += 25
// 	}

// 	// Add 5 points for every two items on receipt
// 	score += 5 * len(receipt.Items) / 2

// 	// Add price * 0.2 if trimmed length of item description % 3 == 0
// 	for _, item := range receipt.Items {
// 		trimmed := strings.TrimSpace(item.ShortDescription)
// 		if len(trimmed)%3 == 0 {
// 			score += int(math.Ceil(item.Price * 0.2))
// 		}
// 	}

// 	// Add 6 points if day in purchase date is odd
// 	dayString := strings.Split(receipt.PurchaseDate, "-")[2]
// 	day, err := strconv.ParseInt(dayString, 10, 8)

// 	if err != nil {
// 		return 0, fmt.Errorf("unable to convert day to int: %v", err)
// 	}

// 	if day%2 == 1 {
// 		score += 6
// 	}

// 	// Add 10 points if time of purchase is after 2pm and before 4pm
// 	time := strings.Split(receipt.PurchaseTime, ":")

// 	if time[0] == "14" && time[1] != "00" || time[0] == "15" {
// 		score += 10
// 	}

// 	return score, nil
// }

func main() {
	receipt, err := deserialzeJSON("examples/target-receipt.json")

	if err != nil {
		fmt.Errorf("Error deserializing JSON: %v", err)
	}

	fmt.Println(receipt)

	// score, err := calculateScore(receipt)

	// if err != nil {
	// 	fmt.Errorf("Unable to calculate score: %v", err)
	// }

	// fmt.Printf("Total score: %v", score)
}

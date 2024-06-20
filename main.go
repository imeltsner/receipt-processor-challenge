package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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

// Calculates points based on total price
func pointsFromTotal(receipt Receipt) (int, error) {
	score := 0

	// Convert total to float
	total, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		return 0, fmt.Errorf("Error converting total to float: %v", err)
	}

	// Adds 50 points if total is round dollar amount
	if int(total*100)%100 == 0 {
		score += 50
	}

	// Add 25 points if total is multiple of 0.25
	if int(total*100)%25 == 0 {
		score += 25
	}

	return score, nil
}

// Calculates points from length of description
func pointsFromDescription(receipt Receipt) int {
	score := 0

	// Add price * 0.2 if trimmed length of item description % 3 == 0
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			score += int(math.Ceil(price * 0.2))
		}
	}

	return score
}

// Calculates the score associated with a given receipt
func calculateScore(receipt Receipt) (int, error) {
	score := 0

	// Adds 1 point for each alphanumeric characters of retailer
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			score += 1
		}
	}

	// Add points from total
	points, _ := pointsFromTotal(receipt)
	score += points

	// Add 5 points for every two items on receipt
	score += 5 * len(receipt.Items) / 2

	// Add score from descriptions
	score += pointsFromDescription(receipt)

	// Add 6 points if day in purchase date is odd
	dayString := strings.Split(receipt.PurchaseDate, "-")[2]
	day, err := strconv.ParseInt(dayString, 10, 8)

	if err != nil {
		return 0, fmt.Errorf("unable to convert day to int: %v", err)
	}

	if day%2 == 1 {
		score += 6
	}

	// Add 10 points if time of purchase is after 2pm and before 4pm
	time := strings.Split(receipt.PurchaseTime, ":")

	if time[0] == "14" && time[1] != "00" || time[0] == "15" {
		score += 10
	}

	return score, nil
}

func main() {
	receipt, err := deserialzeJSON("examples/target-receipt.json")

	if err != nil {
		fmt.Errorf("Error deserializing JSON: %v", err)
	}

	score, err := calculateScore(receipt)

	if err != nil {
		fmt.Errorf("Unable to calculate score: %v", err)
	}

	fmt.Printf("Total score: %v", score)
}

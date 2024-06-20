package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Calculates points based on total price
func pointsFromTotal(receipt Receipt) (int, error) {
	points := 0

	// Convert total to float
	total, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		return 0, fmt.Errorf("error converting total to float: %v", err)
	}

	// Adds 50 points if total is round dollar amount
	if int(total*100)%100 == 0 {
		points += 50
	}

	// Add 25 points if total is multiple of 0.25
	if int(total*100)%25 == 0 {
		points += 25
	}

	return points, nil
}

// Calculates points from length of description
func pointsFromDescription(receipt Receipt) (int, error) {
	points := 0

	// Add price * 0.2 if trimmed length of item description % 3 == 0
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)

			if err != nil {
				return 0, fmt.Errorf("error converting price to float: %v", err)
			}

			points += int(math.Ceil(price * 0.2))
		}
	}

	return points, nil
}

// Calculates points based on if day is odd
func pointsFromDay(receipt Receipt) (int, error) {
	points := 0

	// Parse date string for day
	dayString := strings.Split(receipt.PurchaseDate, "-")[2]
	day, err := strconv.ParseInt(dayString, 10, 8)

	if err != nil {
		return 0, fmt.Errorf("unable to convert day to int: %v", err)
	}

	// Add 6 points if day in purchase date is odd
	if day%2 == 1 {
		points += 6
	}

	return points, nil
}

// Calculates the total points associated with a given receipt
func calculatePoints(receipt Receipt) (int, error) {
	points := 0

	// Adds 1 point for each alphanumeric characters of retailer
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			points += 1
		}
	}

	// Add points from total
	p, err := pointsFromTotal(receipt)
	if err != nil {
		return 0, err
	}
	points += p

	// Add 5 points for every two items on receipt
	points += 5 * (len(receipt.Items) / 2)

	// Add points from description length
	p, err = pointsFromDescription(receipt)
	if err != nil {
		return 0, err
	}
	points += p

	// Add points from day value
	p, err = pointsFromDay(receipt)
	if err != nil {
		return 0, err
	}
	points += p

	// Add 10 points if time of purchase is after 2pm and before 4pm
	time := strings.Split(receipt.PurchaseTime, ":")
	if time[0] == "14" && time[1] != "00" || time[0] == "15" {
		points += 10
	}

	return points, nil
}

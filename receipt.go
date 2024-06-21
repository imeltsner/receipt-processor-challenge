package main

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"required"`
	Total        string `json:"total" validate:"required"`
}

type ID struct {
	ID string `json:"id"`
}

type Points struct {
	Points int `json:"points"`
}

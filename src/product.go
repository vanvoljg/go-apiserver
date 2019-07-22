package main

// Product schema
type Product struct {
	Id          int    `json:"_id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

// Products array
type Products map[int]Product

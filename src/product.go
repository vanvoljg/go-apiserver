package main

// Product schema
type Product struct {
	id          int    `json:"_id"`
	category    string `json:"category"`
	name        string `json:"name"`
	displayName string `json:"display_name"`
	description string `json:"description"`
}

// Products array
type Products map[int]Product

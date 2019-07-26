package main

import "fmt"

// Product schema
type Product struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

func (p Product) String() string {
	return fmt.Sprintf("Product<%v, %v, %v, %v, %v>", p.Id, p.Name, p.Category, p.Description)
}

// // Products array
// type Products map[int]Product

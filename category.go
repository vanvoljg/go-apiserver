package main

// Category schema
type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

// // Categories array
// type Categories map[int]Category

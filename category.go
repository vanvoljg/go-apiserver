package main

import "fmt"

// Category schema
type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

func (c Category) String() string {
	return fmt.Sprintf("Category<%v, %v, %v>", c.Id, c.Name, c.Description)
}

// // Categories array
// type Categories map[int]Category

package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	results := make([]*Product, 0, 10)

	for i, v := range products {
		if v.color == color {
			results = append(results, &products[i])
		}
	}
}

func main() {
	fmt.Println("helo")
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", red, large}
}

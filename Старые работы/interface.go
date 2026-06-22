package main

import "fmt"

type Storable interface {
	GetName() string
	GetCategory() string
	GetPrice() float64
}

type Cream struct {
	Name     string
	Category string
	Price    float64
}

type Perfume struct {
	Name     string
	Category string
	Price    float64
}

type Lipstick struct {
	Name     string
	Category string
	Price    float64
}

func groupByCategory(products []Storable) map[string][]Storable {
	result := make(map[string][]Storable)
	for _, item := range products {
		category := item.GetCategory()
		result[category] = append(result[category], item)
	}
	return result
}

func (c Cream) GetPrice() float64 {
	return c.Price
}

func (c Cream) GetName() string {
	return c.Name
}

func (c Cream) GetCategory() string {
	return c.Category
}

func (p Perfume) GetPrice() float64 {
	return p.Price
}

func (p Perfume) GetName() string {
	return p.Name
}

func (p Perfume) GetCategory() string {
	return p.Category
}

func (l Lipstick) GetPrice() float64 {
	return l.Price
}

func (l Lipstick) GetName() string {
	return l.Name
}

func (l Lipstick) GetCategory() string {
	return l.Category
}

func main() {

	products := []Storable{
		Cream{Name: "Увлажняющий крем", Category: "уход", Price: 1400},
		Perfume{Name: "Dior", Category: "парфюм", Price: 9000},
		Lipstick{Name: "Красная помада", Category: "макияж", Price: 800},
		Cream{Name: "Омолаживающий крем", Category: "уход", Price: 5000},
		Lipstick{Name: "Розовая помада", Category: "макияж", Price: 750},
	}

	index := groupByCategory(products)
	for category, items := range index {
		fmt.Println("Категория:", category)
		for _, item := range items {
			fmt.Printf("  - %s, %.2f руб.\n", item.GetName(), item.GetPrice())
		}
	}
}

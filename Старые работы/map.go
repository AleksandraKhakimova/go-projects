package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type SkinProduct struct {
	Name     string
	Brand    string
	SkinType string
	Price    float64
}

func createIndex(products []SkinProduct) map[string]SkinProduct {
	result := make(map[string]SkinProduct)
	for _, item := range products {
		result[item.Name] = item
	}
	return result
}

func findByName(index map[string]SkinProduct, name string) (SkinProduct, error) {
	product, exists := index[name]
	if !exists { // если НЕ существует
		return SkinProduct{}, errors.New("Продукт не найден")
	}
	return product, nil
}

func main() {

	products := []SkinProduct{
		{Name: "Гель для умывания", Brand: "La Roche", SkinType: "жирная", Price: 750},
		{Name: "Увлажняющий крем", Brand: "Vichy", SkinType: "сухая", Price: 1400},
		{Name: "Тоник для лица", Brand: "La Roche", SkinType: "комбинированная", Price: 650},
		{Name: "Мицеллярная вода", Brand: "Bioderma", SkinType: "чувствительная", Price: 500},
		{Name: "Маска для сухой кожи", Brand: "Vichy", SkinType: "сухая", Price: 1100},
		{Name: "Крем-гель для жирной", Brand: "La Roche", SkinType: "жирная", Price: 950},
	}

	index := createIndex(products)

	var name string
	fmt.Print("Введите название продукта: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()

	product, err := findByName(index, name)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Найден: %s, %s, %.2f руб.\n", product.Name, product.Brand, product.Price)
	}

}

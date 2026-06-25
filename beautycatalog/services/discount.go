// services/discount.go
package services

import "beautycatalog/models"

// Применить скидку ко всем продуктам
func ApplyDiscount(products []models.Product, percent float64) {
	for _, p := range products {
		newPrice := p.GetPrice() * (1 - percent/100)
		p.SetPrice(newPrice)
	}
}
func CountByCategory(products []models.Product) map[string]int {
	result := make(map[string]int)
	for _, p := range products {
		result[p.GetCategory()]++
	}
	return result
}

func AveragePriceByCategory(products []models.Product) map[string]float64 {
	SumCategory := make(map[string]float64)
	for _, p := range products {
		SumCategory[p.GetCategory()] += p.GetPrice()
	}

	counts := CountByCategory(products)
	result := make(map[string]float64)
	for category, sum := range SumCategory {
		result[category] = sum / float64(counts[category])
	}

	return result
}

package storage

import (
	"beautycatalog/models"
	"fmt"
)

type Catalog struct {
	products []models.Product
}

// Конструктор — создаёт пустой каталог
func NewCatalog() *Catalog {
	return &Catalog{}
}

// Добавить продукт
func (c *Catalog) Add(p models.Product) {
	c.products = append(c.products, p)
}

// Получить все продукты
func (c *Catalog) GetAll() []models.Product {
	return c.products
}

func (c *Catalog) GetByCategory(category string) []models.Product {
	var result []models.Product
	for _, p := range c.products {
		if p.GetCategory() == category {
			result = append(result, p)
		}
	}
	return result
}

func (c *Catalog) GetByMaxPrice(maxPrice float64) []models.Product {
	var result []models.Product
	for _, p := range c.products {
		if p.GetPrice() <= maxPrice {
			result = append(result, p)
		}
	}
	return result
}

func (c *Catalog) FindByName(name string) (models.Product, error) {
	for _, p := range c.products {
		if p.GetName() == name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("продукт не найден")
}

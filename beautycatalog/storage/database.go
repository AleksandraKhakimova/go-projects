// storage/database.go
package storage

import (
	"beautycatalog/models"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Database struct {
	db *sql.DB
}

// Открыть базу данных и создать таблицу
func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// Создаём таблицу, если её нет
	query := `CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        category TEXT NOT NULL,
        price REAL NOT NULL,
        product_type TEXT NOT NULL
    );`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

// Добавить продукт
func (d *Database) AddProduct(name, category string, price float64, productType string) error {
	_, err := d.db.Exec(
		"INSERT INTO products (name, category, price, product_type) VALUES (?, ?, ?, ?)",
		name, category, price, productType,
	)
	return err
}

// Получить все продукты
func (d *Database) GetAllProducts() ([]models.Product, error) {
	rows, err := d.db.Query("SELECT name, category, price, product_type FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var name, category, productType string
		var price float64
		rows.Scan(&name, &category, &price, &productType)

		switch productType {
		case "cream":
			products = append(products, &models.Cream{Name: name, Category: category, Price: price})
		case "perfume":
			products = append(products, &models.Perfume{Name: name, Category: category, Price: price})
		case "lipstick":
			products = append(products, &models.Lipstick{Name: name, Category: category, Price: price})
		}
	}
	return products, nil
}
func (d *Database) GetByCategory(category string) ([]models.Product, error) {
	rows, err := d.db.Query(
		"SELECT name, category, price, product_type FROM products WHERE category = ?",
		category,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var name, category, productType string
		var price float64
		rows.Scan(&name, &category, &price, &productType)

		switch productType {
		case "cream":
			products = append(products, &models.Cream{Name: name, Category: category, Price: price})
		case "perfume":
			products = append(products, &models.Perfume{Name: name, Category: category, Price: price})
		case "lipstick":
			products = append(products, &models.Lipstick{Name: name, Category: category, Price: price})
		}
	}
	return products, nil
}

func (d *Database) GetByMaxPrice(maxPrice float64) ([]models.Product, error) {
	rows, err := d.db.Query(
		"SELECT name, category, price, product_type FROM products WHERE price <= ?",
		maxPrice,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var name, category, productType string
		var price float64
		rows.Scan(&name, &category, &price, &productType)

		switch productType {
		case "cream":
			products = append(products, &models.Cream{Name: name, Category: category, Price: price})
		case "perfume":
			products = append(products, &models.Perfume{Name: name, Category: category, Price: price})
		case "lipstick":
			products = append(products, &models.Lipstick{Name: name, Category: category, Price: price})
		}
	}
	return products, nil
}

func (d *Database) FindByName(name string) (models.Product, error) {
	row := d.db.QueryRow(
		"SELECT name, category, price, product_type FROM products WHERE name = ?",
		name,
	)

	var productName, category, productType string
	var price float64
	err := row.Scan(&productName, &category, &price, &productType)
	if err != nil {
		return nil, fmt.Errorf("продукт не найден")
	}

	switch productType {
	case "cream":
		return &models.Cream{Name: productName, Category: category, Price: price}, nil
	case "perfume":
		return &models.Perfume{Name: productName, Category: category, Price: price}, nil
	case "lipstick":
		return &models.Lipstick{Name: productName, Category: category, Price: price}, nil
	}
	return nil, fmt.Errorf("неизвестный тип продукта")
}

// Закрыть базу данных
func (d *Database) Close() {
	d.db.Close()
}

package main

import (
	"beautycatalog/services"
	"beautycatalog/storage"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	db, err := storage.NewDatabase("catalog.db")
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	defer db.Close()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== BeautyCatalog ===")
		fmt.Println("1. Добавить продукт")
		fmt.Println("2. Показать все")
		fmt.Println("3. Фильтр по категории")
		fmt.Println("4. Фильтр по цене")
		fmt.Println("5. Поиск по названию")
		fmt.Println("6. Применить скидку ко всем")
		fmt.Println("7. Статистика")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addProduct(db, scanner)
		case "2":
			ShowAll(db)
		case "3":
			Filter(db, scanner)
		case "4":
			maxPrices(db, scanner)
		case "5":
			searchByName(db, scanner)
		case "6":
			skidka(db, scanner)
		case "7":
			statistika(db)
		case "0":
			fmt.Println("Пока!")
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}

func addProduct(db *storage.Database, scanner *bufio.Scanner) {
	fmt.Print("Тип (cream/perfume/lipstick): ")
	scanner.Scan()
	productType := scanner.Text()

	fmt.Print("Название: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Категория: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Print("Цена: ")
	scanner.Scan()
	price, _ := strconv.ParseFloat(scanner.Text(), 64)

	err := db.AddProduct(name, category, price, productType)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Продукт добавлен!")
}

func ShowAll(db *storage.Database) {
	products, err := db.GetAllProducts()
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	if len(products) == 0 {
		fmt.Println("Каталог пуст")
		return
	}
	for _, p := range products {
		fmt.Printf("%s | Цена: %.2f руб.\n", p.Info(), p.GetPrice())
	}
}
func Filter(db *storage.Database, scanner *bufio.Scanner) {
	fmt.Print("Напишите категорию: ")
	scanner.Scan()
	category := scanner.Text()
	products, err := db.GetByCategory(category)
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("Ничего не найдено")
		return
	}
	for _, p := range products {
		fmt.Printf("%s | Цена: %.2f руб.\n", p.Info(), p.GetPrice())
	}
}

func maxPrices(db *storage.Database, scanner *bufio.Scanner) {
	fmt.Print("Напишите максимальную цену: ")
	scanner.Scan()
	mPrices, _ := strconv.ParseFloat(scanner.Text(), 64)
	products, err := db.GetByMaxPrice(mPrices)
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	if len(products) == 0 {
		fmt.Printf("Нет подходящих под цену")
		return
	}

	for _, p := range products {
		fmt.Printf("%s | Цена: %.2f руб.\n", p.Info(), p.GetPrice())
	}
}

func searchByName(db *storage.Database, scanner *bufio.Scanner) {
	fmt.Print("Напишите название: ")
	scanner.Scan()
	name := scanner.Text()
	products, err := db.FindByName(name)
	if err != nil {
		fmt.Println("Нет подходящих")
		return
	}
	fmt.Printf("%s | Цена: %.2f руб.\n", products.Info(), products.GetPrice())
}

func skidka(db *storage.Database, scanner *bufio.Scanner) {
	fmt.Print("Напишите скидку, которую хотите применить к продуктам: ")
	scanner.Scan()
	percent, _ := strconv.ParseFloat(scanner.Text(), 64)
	allProducts, err := db.GetAllProducts()
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	services.ApplyDiscount(allProducts, percent)
	fmt.Println("Скидка применена ко всем продуктам!")
}

func statistika(db *storage.Database) {
	allProducts, err := db.GetAllProducts()
	if err != nil {
		fmt.Println("Ошибка базы данных:", err)
		return
	}
	if len(allProducts) == 0 {
		fmt.Printf("Нет товаров")
		return
	}

	counts := services.CountByCategory(allProducts)

	fmt.Printf("\nКол-во продуктов по категориям")
	for category, count := range counts {
		fmt.Printf("  %s: %d шт.\n", category, count)
	}

	avgPrices := services.AveragePriceByCategory(allProducts)
	fmt.Println("\nСредняя цена по категориям:")
	for category, avg := range avgPrices {
		fmt.Printf("  %s: %.2f руб.\n", category, avg)
	}
}

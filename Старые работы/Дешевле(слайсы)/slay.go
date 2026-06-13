package main

import (
	"errors"
	"fmt"
)

func inputPrices() ([]float64, error) {

	var count int
	fmt.Printf("Кол-во товаров:")
	fmt.Scan(&count)
	var prices []float64
	for i := 1; i <= count; i++ {
		var price float64
		fmt.Printf("Введите цену товара %d:", i)
		fmt.Scan(&price)
		prices = append(prices, price)
	}
	if len(prices) == 0 {
		return nil, errors.New("список пуст")
	}
	return prices, nil
}

func analyzePrices(prices []float64) (sum float64, avg float64, min float64, max float64, err error) {
	if len(prices) == 0 {
		return 0, 0, 0, 0, errors.New("список пуст")
	}
	min = prices[0]
	max = prices[0]
	for _, price := range prices {

		if price > max {
			max = price
		}
		if price < min {
			min = price
		}
		sum += price
	}
	avg = sum / float64(len(prices))
	return sum, avg, min, max, nil
}

func filterCheap(prices []float64, maxPrice float64) ([]float64, error) {
	var result []float64
	for _, price := range prices {
		if price <= maxPrice {
			result = append(result, price)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("нет товаров дешевле или равных %.2f", maxPrice)
	}
	return result, nil
}

func main() {
	var maxPrice float64

	prices, err := inputPrices()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	sum, avg, min, max, err := analyzePrices(prices)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Сумма:", sum)
	fmt.Println("Средняя:", avg)
	fmt.Println("Минимум:", min)
	fmt.Println("Максимум:", max)

	fmt.Printf("Какой лимит цен:")
	fmt.Scan(&maxPrice)
	result, err := filterCheap(prices, maxPrice)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Товары в лимите:", result)
}

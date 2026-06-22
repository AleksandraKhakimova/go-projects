package main

import (
	"fmt"
)

func searchCreams(query string, ch chan string) {
	creams := []string{"Крем 1", "Крем 2", "Крем 3"}
	for _, name := range creams {
		if name == query {
			ch <- "Найдено в кремах:" + name
			return
		}
	}
	ch <- "В кремах не найдено"
}

func searchPerfumes(query string, ch chan string) {
	perfums := []string{"Dior", "Chanel", "Gucci"}
	for _, name := range perfums {
		if name == query {
			ch <- "Найдено в парфюме " + name
			return
		}
	}
	ch <- "В парфюмах не найдено"
}

func searchLipsticks(query string, ch chan string) {
	lipsticks := []string{"Красная", "Розовая", "Бордовая"}
	for _, name := range lipsticks {
		if name == query {
			ch <- "Найдено в помаде: " + name
			return
		}
	}
	ch <- "В помадах не найдено"
}

func main() {

	var query string
	fmt.Print("Введите название для поиска: ")
	fmt.Scan(&query)

	ch := make(chan string)

	go searchCreams(query, ch)
	go searchPerfumes(query, ch)
	go searchLipsticks(query, ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Конец")
}

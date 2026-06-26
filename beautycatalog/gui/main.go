package main

import (
	"beautycatalog/storage"
	"fmt"
	"strconv"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct{}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 255, G: 105, B: 180, A: 255}
	case theme.ColorNameBackground:
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 128, G: 0, B: 32, A: 255}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 255, G: 255, B: 200, A: 255} // лимонный фон полей
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 200, G: 200, B: 100, A: 255}
	case theme.ColorNameButton:
		return color.NRGBA{R: 255, G: 255, B: 150, A: 255} // лимонный
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func lemonButton(text string, action func()) *widget.Button {
	btn := widget.NewButton(text, action)
	return btn
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return theme.DefaultTheme().Font(style)
	}
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

var db *storage.Database

var (
	pink  = color.NRGBA{R: 255, G: 105, B: 180, A: 255} // розовый
	white = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	green = color.NRGBA{R: 0, G: 200, B: 100, A: 255}
	red   = color.NRGBA{R: 255, G: 80, B: 80, A: 255}
	dark  = color.NRGBA{R: 30, G: 30, B: 30, A: 255}
	light = color.NRGBA{R: 245, G: 245, B: 245, A: 255}
)

func main() {
	var err error
	db, err = storage.NewDatabase("catalog.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("BeautyCatalog")
	w.Resize(fyne.NewSize(700, 500))

	bgImage := canvas.NewImageFromFile("gui/bg.jpg") // или "bg.png"
	bgImage.FillMode = canvas.ImageFillStretch

	menu := mainMenu(w)
	content := container.NewStack(bgImage, menu)

	w.SetContent(content)
	w.ShowAndRun()

}

func coloredButton(text string, bg color.NRGBA, action func()) *widget.Button {
	btn := widget.NewButton(text, action)
	return btn
}

// Кнопка "Назад в меню"
func backBtn(w fyne.Window) *widget.Button {
	btn := widget.NewButton("← Назад в меню", func() {
		w.SetContent(mainMenu(w))
	})
	return btn
}

// Главное меню
func mainMenu(w fyne.Window) fyne.CanvasObject {
	menu := container.NewVBox(
		widget.NewLabelWithStyle("=== BeautyCatalog ===", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		coloredButton("1. Показать все продукты", pink, func() { w.SetContent(showAllScreen(w)) }),
		coloredButton("2. Добавить продукт", pink, func() {
			w.SetContent(addProductScreen(w))
		}),

		coloredButton("3. Фильтр по категории", pink, func() {
			w.SetContent(filterByCategoryScreen(w))
		}),
		coloredButton("4. Фильтр по цене", pink, func() {
			w.SetContent(filterByPriceScreen(w))
		}),
		coloredButton("5. Поиск по названию", pink, func() {
			w.SetContent(searchScreen(w))
		}),
		coloredButton("6. Применить скидку", pink, func() {
			w.SetContent(discountScreen(w))
		}),
		coloredButton("7. Статистика", pink, func() {
			w.SetContent(statsScreen(w))
		}),
		coloredButton("0. Выход", red, func() {
			w.Close()
		}),
	)
	return backgroundScreen(w, "bg.jpg", container.NewCenter(menu))
}

// 1. Показать все
func showAllScreen(w fyne.Window) fyne.CanvasObject {
	products, _ := db.GetAllProducts()
	list := widget.NewList(
		func() int { return len(products) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			p := products[id]
			obj.(*widget.Label).SetText(fmt.Sprintf("%s | Цена: %.2f руб.", p.Info(), p.GetPrice()))
		},
	)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, list))
}

// 2. Добавить продукт
func addProductScreen(w fyne.Window) fyne.CanvasObject {
	typeSelect := widget.NewSelect([]string{"cream", "perfume", "lipstick"}, func(s string) {})
	typeSelect.SetSelected("cream")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название")
	categoryEntry := widget.NewEntry()
	categoryEntry.SetPlaceHolder("Категория")
	priceEntry := widget.NewEntry()
	priceEntry.SetPlaceHolder("Цена")

	statusLabel := widget.NewLabel("")
	statusLabel.TextStyle = fyne.TextStyle{Bold: true}

	addBtn := widget.NewButton("Добавить", func() {
		price, err := strconv.ParseFloat(priceEntry.Text, 64)
		if err != nil {
			statusLabel.SetText("Ошибка: введите число в поле цены")
			return
		}

		db.AddProduct(nameEntry.Text, categoryEntry.Text, price, typeSelect.Selected)
		statusLabel.SetText("Продукт добавлен!")
		nameEntry.SetText("")
		categoryEntry.SetText("")
		priceEntry.SetText("")
	})

	addBtn.Importance = widget.HighImportance

	form := container.NewVBox(
		widget.NewLabel("Тип:"),
		typeSelect,
		nameEntry,
		categoryEntry,
		priceEntry,
		addBtn,
		statusLabel,
	)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, container.NewCenter(form)))

}

// 3. Фильтр по категории
func filterByCategoryScreen(w fyne.Window) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Категория")
	resultLabel := widget.NewLabel("")

	searchBtn := widget.NewButton("Показать", func() {
		products, _ := db.GetByCategory(entry.Text)
		if len(products) == 0 {
			resultLabel.SetText("Ничего не найдено")
			return
		}
		text := ""
		for _, p := range products {
			text += fmt.Sprintf("%s | Цена: %.2f руб.\n", p.Info(), p.GetPrice())
		}
		resultLabel.SetText(text)
	})

	content := container.NewVBox(entry, searchBtn, resultLabel)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, content))

}

// 4. Фильтр по цене
func filterByPriceScreen(w fyne.Window) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Максимальная цена")
	resultLabel := widget.NewLabel("")

	searchBtn := widget.NewButton("Показать", func() {
		maxPrice, _ := strconv.ParseFloat(entry.Text, 64)
		products, _ := db.GetByMaxPrice(maxPrice)
		if len(products) == 0 {
			resultLabel.SetText("Ничего не найдено")
			return
		}
		text := ""
		for _, p := range products {
			text += fmt.Sprintf("%s | Цена: %.2f руб.\n", p.Info(), p.GetPrice())
		}
		resultLabel.SetText(text)
	})

	content := container.NewVBox(entry, searchBtn, resultLabel)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, content))
}

// 5. Поиск по названию
func searchScreen(w fyne.Window) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Название")
	resultLabel := widget.NewLabel("")

	searchBtn := widget.NewButton("Найти", func() {
		p, err := db.FindByName(entry.Text)
		if err != nil {
			resultLabel.SetText("Продукт не найден")
			return
		}
		resultLabel.SetText(fmt.Sprintf("%s | Цена: %.2f руб.", p.Info(), p.GetPrice()))
	})

	content := container.NewVBox(entry, searchBtn, resultLabel)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, content))

}

// 6. Скидка
func discountScreen(w fyne.Window) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Процент скидки")
	resultLabel := widget.NewLabel("")

	applyBtn := widget.NewButton("Применить", func() {
		percent, _ := strconv.ParseFloat(entry.Text, 64)
		products, _ := db.GetAllProducts()
		for _, p := range products {
			newPrice := p.GetPrice() * (1 - percent/100)
			p.SetPrice(newPrice)
		}
		resultLabel.SetText(fmt.Sprintf("Скидка %.0f%% применена!", percent))
	})

	content := container.NewVBox(entry, applyBtn, resultLabel)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, content))
}

// 7. Статистика
func statsScreen(w fyne.Window) fyne.CanvasObject {
	products, _ := db.GetAllProducts()
	text := ""

	counts := make(map[string]int)
	sums := make(map[string]float64)
	for _, p := range products {
		counts[p.GetCategory()]++
		sums[p.GetCategory()] += p.GetPrice()
	}

	text += "=== Количество ===\n"
	for cat, cnt := range counts {
		text += fmt.Sprintf("%s: %d шт.\n", cat, cnt)
	}

	text += "\n=== Средняя цена ===\n"
	for cat, sum := range sums {
		text += fmt.Sprintf("%s: %.2f руб.\n", cat, sum/float64(counts[cat]))
	}

	label := widget.NewLabel(text)
	return backgroundScreen(w, "bg_common.jpg", container.NewBorder(nil, backBtn(w), nil, nil, label))
}

func backgroundScreen(w fyne.Window, bgName string, content fyne.CanvasObject) fyne.CanvasObject {
	bgImage := canvas.NewImageFromFile("gui/" + bgName)
	bgImage.FillMode = canvas.ImageFillStretch
	return container.NewStack(bgImage, content)
}

package models

type Product interface {
	GetName() string
	GetCategory() string
	GetPrice() float64
	SetPrice(price float64)
	Info() string
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

func (c Cream) GetName() string {
	return c.Name
}
func (c Cream) GetCategory() string {
	return c.Category
}
func (c Cream) GetPrice() float64 {
	return c.Price
}
func (c *Cream) SetPrice(p float64) {
	c.Price = p
}
func (c Cream) Info() string {
	return "Крем: " + c.Name + " | Категория: " + c.Category
}

func (f Perfume) GetName() string {
	return f.Name
}
func (f Perfume) GetCategory() string {
	return f.Category
}
func (f Perfume) GetPrice() float64 {
	return f.Price
}
func (f *Perfume) SetPrice(p float64) {
	f.Price = p
}
func (f Perfume) Info() string {
	return "Парфюм: " + f.Name + " | Категория: " + f.Category
}

func (l Lipstick) GetName() string {
	return l.Name
}
func (l Lipstick) GetCategory() string {
	return l.Category
}
func (l Lipstick) GetPrice() float64 {
	return l.Price
}
func (l *Lipstick) SetPrice(p float64) {
	l.Price = p
}
func (l Lipstick) Info() string {
	return "Крем: " + l.Name + " | Категория: " + l.Category
}

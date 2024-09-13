package main

import (
	"fmt"
	"math"
)

type Person struct {
	firstName  string
	lastName   string
	middleName string
	age        int
}

func (p Person) getPersonInfo() {
	fmt.Println("Имя: ", p.firstName)
	fmt.Println("Фамилия: ", p.lastName)
	fmt.Println("Отчество: ", p.middleName)
	fmt.Println("Возраст: ", p.age)
}

func (p *Person) birthday() {
	p.age++
}

type Circle struct {
	radius float64
}

// Фигуры реализуют интерфейс, если у них будут метод Area
type Shape interface {
	Area() float64
}

type Rectangle struct {
	a float64
	b float64
}

// Метод Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.a * r.b
}

// Метод Area для Circle
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func callShapeArea(shapes []Shape) {

	for _, v := range shapes {
		fmt.Println(v.Area())
	}
}

func main() {
	//region #1 Exercise
	p1 := Person{
		"Михаил",
		"Зубенко",
		"Петрович",
		228,
	}
	p1.getPersonInfo()

	//region #2 Exercise
	p1.birthday()
	p1.getPersonInfo()

	//region #3-4 Exercise
	r1 := Rectangle{a: 2, b: 3}
	fmt.Println(r1.Area())

	c1 := Circle{radius: 3}
	fmt.Println(c1.Area())

	r2 := Rectangle{a: 5, b: 2}

	//region #5 Exercise
	Shapes := []Shape{c1, r1, r2}

	fmt.Println("#5 Exercise")

	callShapeArea(Shapes)

	//region #6 Exercise
	var b1 Stringer = &Book{
		id:    1,
		title: "Капитал",
		price: 228}
	b1.getInfoBook()
	b1.setNewPriceInfoBook(1337)
	b1.getInfoBook()
}

type Stringer interface {
	getInfoBook()
	setNewPriceInfoBook(int)
}

type Book struct {
	id    int
	title string
	price int
}

func (b Book) getInfoBook() {
	fmt.Println("Айди: ", b.id)
	fmt.Println("Название: ", b.title)
	fmt.Println("Цена: ", b.price)
}

func (b *Book) setNewPriceInfoBook(newPrice int) {
	b.price = newPrice
}

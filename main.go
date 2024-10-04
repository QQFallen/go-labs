package main

import (
	"fmt"
	"math"
)

func main() {
	pers := Person{age: 12, name: "Alasgvfl"}
	pers.getpers()
	pers = pers.birhday()
	pers.getpers()

	circ := Circle{radius: 5.3}
	rect := Rectangle{width: 5, heigth: 6}

	fmt.Println("\nВывод по одиночке(без среза):")
	PrintArea(circ)
	PrintArea(rect)

	fmt.Println("\nВывод с помощью среза:")
	shapes := []Shape{rect, circ}
	PrintAreas(shapes)

	book := Book{title: "...", author: "...", age: 2024}
	fmt.Println(book)

}

// Функция для вывода среза shape
func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Println("Площадь фигуры ->", shape.Area())
	}
}

// функции для Person
func (r Person) getpers() {
	fmt.Printf("%s - Возраст -> %d\n", r.name, r.age)
}
func (r Person) birhday() Person {
	r.age += 1
	fmt.Println("Поздравляем с днём рождения!")
	return r
}

// Создание структур и интерфейсов
type Person struct {
	age  int
	name string
}

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	heigth float64
}
type Circle struct {
	radius float64
}

type Book struct {
	age    int
	title  string
	author string
}

type stringer interface{ String() string }
// функция для Book
func (b Book) String() string {
	return fmt.Sprintf("\nКнига:%s\nАвтор:%s\nГод написания:%d", b.title, b.author, b.age)
}

// функции площади для интерфейса
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r Rectangle) Area() float64 {
	return r.heigth * r.width
}
func PrintArea(x Shape) {
	fmt.Println("Площадь фигуры ->", x.Area())
}

package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Посетитель позволяет добавлять поведение в структуру без ее изменения.
Посетитель предлагает разместить новое поведение в отдельном классе,
вместо того чтобы множить его сразу в нескольких классах.
Объекты будут передаваться в методы посетителя, методы, скорее всего будут разные и придется проверять
тип объекта, чтобы выбрать метод
+ можем не перегружать класс объекта
+ объеденяем родственные опреации

Шаги:
1. Определяем интерфейс посетителя с функциями для каждого типа

2. Для каждого типа пишем метод, который будет принимать Посетителя

*/

import (
	"fmt"
	"math"
)

// Интерфейс Посетителя
type Visitor interface {
	visit(Shape)
}

// Интерфейс для всех геометрических фигур
type Shape interface {
	accept(v Visitor)
}

// Структура типа
type Square struct {
	side float64
}

func (s Square) accept(v Visitor) {
	v.visit(s)
}

// Структура типа
type Circle struct {
	radius float64
}

func (c Circle) accept(v Visitor) {
	v.visit(c)
}

// Сам Посетитель
type AreaCalculator struct {
	area float64
}

// Используем утверждение типа
func (a AreaCalculator) visit(s Shape) {
	switch s := s.(type) {
	case Square:
		a.visitForSquare(s)
	case Circle:
		a.visitForCircle(s)
	default:
		fmt.Println("Unsupported shape")
	}
}

func (a *AreaCalculator) visitForSquare(s Square) {
	a.area = s.side * s.side
	fmt.Println("Area of square = ", a.area)
}

func (a *AreaCalculator) visitForCircle(c Circle) {
	a.area = c.radius * c.radius * math.Pi
	fmt.Printf("Area of circle: %.2f\n", a.area)
}

// func main() {
// 	ac := AreaCalculator{}
// 	c := Circle{radius: 3.0}

// 	c.accept(ac)
// 	s := Square{side: 2}
// 	s.accept(ac)
// }

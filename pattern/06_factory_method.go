package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Используем когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.

Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.

Благодаря этому, код производства можно расширять, не трогая основной. Так, чтобы добавить поддержку нового продукта, вам нужно создать новый подкласс и определить в нём фабричный метод, возвращая оттуда экземпляр нового продукта.

*/


import (
	"fmt"
)

// Определение констант для типов хранимых объектов.
const (
	ServType     string = "serv" // Тип для серверной аппаратуры.
	NotebookType string = "notebook" // Тип для ноутбука.
)

// StoreObject интерфейс, определяющий методы для получения типа объекта и вывода его деталей.
type StoreObject interface {
	GetType() string
	PrintDetails()
}

// Notebook структура для описания ноутбука.
type Notebook struct {
	Display  string // Дисплей ноутбука.
	Keyboard string // Клавиатура ноутбука.
	Trackpad string // Трекпад ноутбука.
}

// Serv структура для описания серверной аппаратуры.
type Serv struct {
	CPU    string // Процессор сервера.
	Memory int   // Оперативная память сервера.
}

// NewServ создает новый экземпляр структуры Serv.
func NewServ() Serv {
	return Serv{
		CPU:    "Intel", // Пример процессора.
		Memory: 256,    // Пример объема оперативной памяти.
	}
}

// NewNotebook создает новый экземпляр структуры Notebook.
func NewNotebook() Notebook {
	return Notebook{
		Display:  "HP", // Пример производителя дисплея.
		Keyboard: "FullSize", // Пример типа клавиатуры.
		Trackpad: "Crap", // Пример состояния трекпада.
	}
}

// GetType возвращает строку, указывающую тип объекта.
func (n Notebook) GetType() string {
	return "notebook"
}

// PrintDetails выводит детали объекта Notebook.
func (n Notebook) PrintDetails() {
	fmt.Printf("Display %s, Keyboard %s, Trackpad %s\n", n.Display, n.Keyboard, n.Trackpad)
}

// GetType возвращает строку, указывающую тип объекта.
func (s Serv) GetType() string {
	return "serv"
}

// PrintDetails выводит детали объекта Serv.
func (s Serv) PrintDetails() {
	fmt.Printf("CPU %s, Mem %d\n", s.CPU, s.Memory)
}

// New создает новый объект на основе переданного типа.
func New(typeName string) StoreObject {
	switch typeName {
	default:
		fmt.Printf("Несуществующий тип %s\n", typeName)
		return nil
	case ServType:
		return NewServ()
	case NotebookType:
		return NewNotebook()
	}
}

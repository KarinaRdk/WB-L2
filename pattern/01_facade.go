package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

Элементы: клиент, фасад, система

Фасад — это структурный паттерн, который предоставляет простой (но урезанный)
интерфейс к сложной системе объектов, библиотеке или фреймворку.

Фасад позволяет клиенту не взаимодействовать клиенту с системой напрямую, и переадресует
запрос необходимым частям системы.
Система о фасаде не знает.
Фасад может не давать доступ ко все фичам, но только к тем, которые необходимы клиенту =>
может быть полезен, когда мы используем большую библиотеку, но нам нужна только часть ее функциональности

Отличия от адептера:
Фасад задаёт новый интерфейс, Адаптер повторно использует старый.
Адаптер оборачивает только один класс, а фасад оборачивает целую подсистему.
*/

import "fmt"

// Facade интерфейс для фасада
type Facade interface {
	ProcessData(data string)
}

// MyFacade реализация интерфейса Facade
type MyFacade struct {
	Logger   Logger
	Database Database
}

func NewMyFacade(logger Logger, db Database) *MyFacade {
	return &MyFacade{
		Logger:   logger,
		Database: db,
	}
}

func (f *MyFacade) ProcessData(data string) {
	f.Logger.Log("Processing data...")
	err := f.Database.SaveData(data)
	if err != nil {
		f.Logger.Log("Error saving data")
	} else {
		f.Logger.Log("Data saved successfully")
	}
}

// ConsoleLogger реализация интерфейса Logger
type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(message string) {
	fmt.Println("Log:", message)
}

// SimpleDB реализация интерфейса Database
type SimpleDB struct{}

func (db *SimpleDB) SaveData(data string) error {
	fmt.Println("Save data:", data)
	return nil
}

// Logger интерфейс для логгера
type Logger interface {
	Log(message string)
}

// Database интерфейс для базы данных
type Database interface {
	SaveData(data string) error
}

func main() {
	logger := &ConsoleLogger{}
	db := &SimpleDB{}

	facade := NewMyFacade(logger, db)
	facade.ProcessData("Some data")

	log.Println("Program finished")

}

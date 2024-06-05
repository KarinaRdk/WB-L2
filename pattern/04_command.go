package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
а также поддерживать отмену операций.

Структура:

Клиент создаёт объекты конкретных команд, передавая в них все необходимые параметры
После этого клиент связывает объекты отправителей с созданными командами.

Отправитель хранит ссылку на объект команды и обращается к нему, когда нужно выполнить какое-то действие.
Отправитель работает с командами только через их общий интерфейс. Он не знает, какую конкретно команду
использует, так как получает готовый объект команды от клиента.

Команда описывает общий для всех конкретных команд интерфейс. Обычно здесь описан всего один метод для запуска команды.

Конкретные команды реализуют различные запросы, следуя общему интерфейсу команд. Обычно команда не делает всю работу
самостоятельно, а лишь передаёт вызов получателю, которым является один из объектов бизнес-логики.

Получатель содержит бизнес-логику программы. В этой роли может выступать практически любой объект.

Применяется:
- чтобы параметризовать объекты выполняемым действием.
- чтобы ставить операции в очередь, выполнять их по расписанию или передавать по сети.
- для операции отмены

+ Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
+ Позволяет реализовать простую отмену и повтор операций.
+ Позволяет реализовать отложенный запуск операций.
+ Реализует принцип открытости/закрытости.
- Усложняет код программы из-за введения множества дополнительных классов.


Посетитель можно рассматривать как расширенный аналог Команды, который способен работать сразу с несколькими видами получателей.


Шаги:
Создайте общий интерфейс команд и определите в нём метод запуска.
создайте классы конкретных команд. В каждом классе должно быть поле
   для хранения ссылки на один или несколько объектов-получателей,
Кроме этого, команда должна иметь поля для хранения параметров, которые
нужны при вызове методов получателя. Значения полей команда должна получать через конструктор.
реализуйте основной метод команды, вызывая в нём те или иные методы получателя.

Добавьте в классы отправителей поля для хранения команд.
Обычно отправители принимают команды через конструктор либо через сеттер

Порядок инициализации объектов:

    Создаём объекты получателей.
    Создаём объекты команд, связав их с получателями.
    Создаём объекты отправителей, связав их с командами.


*/

import (
	"fmt"
)
// Интерфейс команды
type Command interface {
	execute()
}

// Конкретная команда
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Отправитель
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Интерфейс получателя
type Device interface {
	on()
	off()
}

// Конкретный получатель
type Tv struct {
	isRunning bool
}

// Методы получателя
func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	// Создаем объект получателя
	tv := &Tv{}
	// Создаем команды и связываем с получателем
	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}
	// Создаем отправителя и связываем с командой
	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}

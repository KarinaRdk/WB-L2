
/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/
package main

import (
	"fmt"
	"time"
)

// or объединяет несколько каналов в один, закрывая результирующий канал при получении данных из любого из входящих каналов.
func or(channels...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	for _, val := range channels {
		go func(ch <-chan interface{}) {
			<-ch // Ожидание получения данных из канала.
			close(res) // Закрытие результирующего канала после получения данных.
		}(val)
	}
	return res
}

func main() {
	// Функция sig создает канал, который закрывается через указанный промежуток времени.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c) // Закрытие канала после завершения таймаута.
			time.Sleep(after) // Ожидание указанного времени перед закрытием канала.
		}()
		return c
	}

	start := time.Now() // Запись времени начала выполнения программы.

	// or используется для объединения нескольких сигналов (каналов), каждый из которых закрывается через разное время.
	<-or(
		sig(2*time.Hour), // Канал, закрывающийся через 2 часа.
		sig(5*time.Minute), // Канал, закрывающийся через 5 минут.
		sig(1*time.Second), // Канал, закрывающийся через 1 секунду.
		sig(1*time.Hour), // Канал, закрывающийся через 1 час.
		sig(1*time.Minute), // Канал, закрывающийся через 1 минуту.
	)

	fmt.Printf("done after %v", time.Since(start)) // Вывод времени, прошедшего с начала выполнения до закрытия всех каналов.
}

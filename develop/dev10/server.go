package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// myHandle обрабатывает входящие соединения от клиентов, принимая сообщения и отвечая на них.
func myHandle(conn net.Conn) {
	for {
		// Чтение сообщения от клиента до символа новой строки ('\n').
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Обработка сообщения в зависимости от его содержимого.
		switch message {
		case "time is up":
			// Завершение обработки соединения, если клиент отправил "time is up".
			return
		default:
			// Вывод полученного сообщения.
			fmt.Print("Message Received:", message)
			// Преобразование сообщения в верхний регистр.
			newMessage := strings.ToUpper(message)
			// Отправка обратного сообщения клиенту с текущим временем и преобразованным сообщением.
			_, err := conn.Write([]byte(time.Now().String() + " " + newMessage + "\n"))
			if err!= nil {
				log.Print(err)
			}
		}
	}
}

func main() {
	// Начало прослушивания TCP-соединений на порту 3000.
	ln, _ := net.Listen("tcp", "localhost:3000")
	// Принятие первого входящего соединения.
	conn, _ := ln.Accept()

	// Обработка соединения.
	myHandle(conn)

	// Закрытие сокета прослушивания.
	err := ln.Close()
	if err!= nil {
		log.Print(err)
	}
}

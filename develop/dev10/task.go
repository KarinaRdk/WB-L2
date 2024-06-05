package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// Client представляет клиентское соединение к серверу.
type Client struct {
	host     string // Хост сервера.
	port     string // Порт сервера.
	connTime time.Duration // Время подключения к серверу.
}

// NewClient создает новый экземпляр Client с заданными параметрами.
func NewClient(h, p, cTime string) *Client {
	t, err := strconv.Atoi(cTime)
	if err!= nil {
		log.Fatal(err)
	}
	return &Client{
		host:     h,
		port:     p,
		connTime: time.Duration(t) * time.Second,
	}
}

// dial устанавливает соединение с сервером и обрабатывает взаимодействие между клиентом и сервером.
func dial(client *Client, ctx context.Context, cancel context.CancelFunc) {
	conn, err := net.DialTimeout("tcp", client.host+":"+client.port, client.connTime)
	if err!= nil {
		log.Fatal(err)
	}
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			// Отправляет сообщение серверу, когда время истекло.
			_, err := fmt.Fprintf(conn, "time is up")
			if err!= nil {
				log.Print(err)
			}
			log.Println("time is up...")
			err = conn.Close()
			if err!= nil {
				log.Print(err)
			}
			return
		default:
			// Чтение ввода пользователя и отправка его серверу.
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("message: ")
			text, err := rd.ReadString('\n')
			if err!= nil {
				log.Print("read error...")
			}
			_, err = fmt.Fprintf(conn, text+"\n")
			if err!= nil {
				log.Print(err)
			}
			// Чтение ответа сервера.
			fb, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Println("from server :" + fb)
		}
	}
}

func main() {
	// Определение флага для указания времени работы с сервером.
	t := flag.String("timeout", "10", "время на работу с сервером")
	flag.Parse()

	// Проверка корректности введенных аргументов.
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("should be host and port")
	}
	port := args[1]
	host := args[0]

	// Создание нового клиента с заданными параметрами.
	client := NewClient(host, port, *t)

	// Создание контекста с ограниченным временем.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, client.connTime)

	// Установление соединения и взаимодействие с сервером.
	dial(client, ctx, cancel)

	fmt.Println("Connection closed...")
}

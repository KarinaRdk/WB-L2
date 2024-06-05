
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTellTime(t *testing.T) {
	// Создаем мок HTTP сервер, который возвращает фиктивное время
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.Write([]byte(`time=2024-05-08T12:34:56Z`))
	}))
	defer server.Close()

	// Используем IP адрес мок сервера
	tellTime(server.Listener.Addr().String())

	// Проверяем, что функция печатает ожидаемое время
	// Здесь вы можете добавить проверку вывода функции, например, сравнить его с ожидаемым значением
}


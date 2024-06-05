package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/pborman/getopt"
)

// getFile считывает содержимое файла, возможно, исключая повторяющиеся строки в зависимости от флага uniqueRequred.
func getFile(path string, uniqueRequred bool) ([]string, error) {
	file, err := os.Open(path)
	set := make(map[string]struct{}) // Используется map для отслеживания уникальности строк.
	if err!= nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	data := make([]string, 0) // Срез для хранения прочитанных строк.
	for scanner.Scan() {
		if uniqueRequred {
			if _, ok := set[scanner.Text()];!ok { // Проверяем, была ли строка уже добавлена.
				set[scanner.Text()] = struct{}{} // Добавляем строку в map, если она еще не была добавлена.
			} else {
				continue // Пропускаем строку, если она уже была добавлена.
			}
		}
		data = append(data, scanner.Text()) // Добавляем строку в срез.
	}
	return data, nil
}

// toSort сортирует строки в срезе. Если флаг n установлен, сортировка происходит по числовым значениям строк.
func toSort(data []string, n bool) []string {
	if n {
		sort.Slice(data, func(i, j int) bool {
			vi, _ := strconv.Atoi(data[i]) // Преобразование строки в число для сравнения.
			vj, _ := strconv.Atoi(data[j]) // Преобразование строки в число для сравнения.
			return vi < vj // Сортировка по возрастанию числового значения.
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j] // Сортировка по возрастанию строки.
		})
	}
	return data
}

func main() {
	filename := getopt.String('f', "", "файл") // Получение имени файла из командной строки.
	n := *getopt.Bool('n', "сортировка по числовому значению") // Получение флага сортировки по числовому значению.
	r := *getopt.Bool('r', "сортировка в обратном порядке") // Получение флага сортировки в обратном порядке.
	u := *getopt.Bool('u', "не выводить повторяющиеся строки") // Получение флага исключения повторяющихся строк.
	getopt.Parse() // Парсинг командной строки.

	file, err := getFile(*filename, u) // Чтение файла и возможное исключение повторяющихся строк.
	if err!= nil {
		panic(err) // Завершение программы в случае ошибки чтения файла.
	}

	file = toSort(file, n) // Сортировка строк в файле.

	if r {
		for i := len(file) - 1; i > 0; i-- { // Вывод строк в обратном порядке.
			fmt.Println(file[i])
		}
	} else {
		for _, value := range file { // Вывод строк в прямом порядке.
			fmt.Println(value)
		}
	}
}

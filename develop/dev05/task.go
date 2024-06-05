/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/pborman/getopt"
)

// openFile открывает файл по указанному пути и считывает его содержимое в виде списка строк.
func openFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err!= nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

// getExpression компилирует регулярное выражение с возможностью игнорировать регистр.
func getExpression(pattern string, ignore bool) (*regexp.Regexp, error) {
	ignorePrefix := ""
	if ignore {
		ignorePrefix = "(?i)" // Игнорирование регистра в регулярном выражении.
	}
	compiledExpession, err := regexp.Compile(ignorePrefix + pattern)
	if err!= nil {
		return nil, err
	}
	return compiledExpession, nil
}

// getNumberOfIntersections подсчитывает количество строк в файле, которые соответствуют заданному регулярному выражению.
func getNumberOfIntersections(file []string, expression *regexp.Regexp) int {
	result := 0
	for _, str := range file {
		match := expression.Match([]byte(str))
		if match {
			result++
		}
	}
	return result
}

// reg проходит по каждой строке файла и выводит контекст вокруг совпадений с регулярным выражением.
func reg(file []string, expression *regexp.Regexp, after, before int, number, invert bool) {
	for i, str := range file {
		match := expression.Match([]byte(str))
		if invert &&!match {
			echo(file, i, after, before, number)
		} else if!invert && match {
			echo(file, i, after, before, number)
		}
	}
}

// echo выводит контекст вокруг найденного совпадения с регулярным выражением.
func echo(file []string, i, after, before int, number bool) {
	startPoint := 0
	endPoint := len(file)
	if i-after > 0 {
		startPoint = i - after
	}
	if i+before < len(file) {
		endPoint = i + before
	}
	if endPoint!= len(file) {
		endPoint += 1
	}
	fmt.Println("------------------------")
	for line := startPoint; line < endPoint; line++ {
		if number {
			fmt.Printf("%d: ", line+1)
		}
		fmt.Printf("%s\n", file[line])
	}
	fmt.Println("------------------------")
}

func main() {
	pattern := getopt.String('e', "", "паттерн")
	path := getopt.String('f', "", "файл")
	after := getopt.IntLong("after", 'A', 0, "вывод N строк после совпадения")
	before := getopt.IntLong("before", 'B', 0, "вывод N строк до совпадения")
	inTheMiddle := getopt.IntLong("context", 'C', 0, "вывод N строк в районе совпадения")
	count := getopt.Bool('c', "вывести количество строк с совпадением")
	ignore := getopt.Bool('i', "игнорировать различия регистра")
	invert := getopt.Bool('v', "инвертировать вывод")
	number := getopt.Bool('n', "напечатать номер строки")

	getopt.Parse()

	file, err := openFile(*path)
	if err!= nil {
		panic(err)
	}
	expression, err := getExpression(*pattern, *ignore)
	if err!= nil {
		panic(err)
	}

	if *count {
		result := getNumberOfIntersections(file, expression)
		if *invert {
			result = len(file) - result
		}
		fmt.Println(result)
	} else {
		if *after == 0 && *before == 0 && *inTheMiddle!= 0 {
			*after = *inTheMiddle / 2
			*before = *inTheMiddle / 2
		}
		reg(file, expression, *after, *before, *number, *invert)
	}
}

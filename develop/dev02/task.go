package main
import (
	"fmt"
	"unicode"
	
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	a := "1"
	fmt.Println(unpack(a))

}
/*
идем по строке двумя индексами
если правый на цифре и левый на букве - левый кладем нужное количество раз
сдигаем оба 
если правый и левый на цифре - возвращаем ошибку
если правй на букве и левый на букве - записываем левый 
сдвигаем оба 

*/
func unpack(input string) (string, error) {
	if len(input) < 1 {
		return input, nil
	}
	s := []rune(input)
	var answer[]rune
	j := 1
	i := 0
	for ; i < len(s)-1 && j < len(s); i++ {
		fmt.Println(string(s[i]), string(s[j]))
		if unicode.IsDigit(s[j]) {
			if !unicode.IsDigit(s[i]) {
				for n := 0; n < int(s[j]-'0'); n++ {
					answer = append(answer, s[i])
				}
			} else {
				return input, fmt.Errorf("invalid input")
			}
		} else { 
			if !unicode.IsDigit(s[i]) {
				answer = append(answer, s[i])
			}
		}
		j++

	}
	if !unicode.IsDigit(s[i]) {
		answer = append(answer, s[i])
	}
	if len(answer) == 0 {
		return input, fmt.Errorf("invalid input")
	}
	
		return string(answer), nil

}
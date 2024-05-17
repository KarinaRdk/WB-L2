Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++  
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
2
1

в первой функции мы увеличили значение х перед выходом из функции, во втором случае нет, поскольку в первом у нас соблюдается 2 условия: у окружающей функции поименнованное возвращаемое значение, defer откладывает исполнение анонимной функции

https://go.dev/ref/spec#Defer_statements

 if the deferred function is a function literal (anonymous function) and the surrounding function has named result parameters that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. 

```

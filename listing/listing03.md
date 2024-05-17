Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false

динамическое значение интерфейса будет рано нил, но значение интерфейса - нет

Значение интерфейса равно нил, только если и значение и динамический тип равны нил. Foo() возвращает [nil, *os.PathError], что не равно [nil, nil].

https://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/
```

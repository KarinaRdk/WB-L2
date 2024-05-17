Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
.error
чтобы вывести ок сравнение должно быть с нил, приведённому к типу указателя на кастомную ошибку
if err !=(*customError)(nil)

Значение интерфейса равно нил, только если и значение и динамический тип равны нил. test() возвращает [nil, *customError], что не равно [nil, nil].

https://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/
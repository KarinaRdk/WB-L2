Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
будут выведены все значения, переданные в оба канала, затем будут печататься нулевые значения типа, поскольку будет происходить чтение из закрытого канала

исправленный merge()  будет прекращать чтоение после закрытия канала и выводить только значения, записанные в канал:

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for a != nil || b!= nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					break
				}
				c <- v
			case v, ok := <-b: 
			if !ok {
				b = nil
				break
			}
				c <- v
			}
		}
		close(c)
	}()
	return c
}

```

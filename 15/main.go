/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить?
Приведите корректный пример реализации.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createHugeString(length int) string {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func someFunc(i int) string {
	v := createHugeString(i)
	return v
}

func main() {
	fmt.Println(someFunc(100))
	fmt.Println(someFunc(200))
	fmt.Println(someFunc(1 << 10))
	fmt.Println(someFunc(10))
}

/*
	Я убрала глобальную переменную. Глобальные переменные являются антипаттерном и их использование
	чаще всего не оправдано.
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}
	Создавать большую строку, а затем брать фиксированный срез в 100 символов, -
	это бессмысленная трата ресурсов.
	Я дополнила someFunc тем, что теперь она принимает число - желаемый размер строки
	и возвращает эту же строку.
	Теперь я могу вызвать и записать результат выполнения этой функции, передавая в нее
	длину строки. Благодаря этому мне не нужно создавать срез.
*/

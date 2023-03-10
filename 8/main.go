/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

package main

import (
	"fmt"
)

func main() {
	var number int64 = 123
	var i uint = 5 // номер бита, который будем менять, начиная с нулевого

	fmt.Printf("%b\n", number) // посмотрим на двоичное представление числа

	number |= 1 << i // устанавливаем i-й бит в 1, OR установит единицу; дизъюнкция
	fmt.Printf("%d-й бит установлен в 1: %d\nДвоичное представление: %0b\n", i, number, number)

	number &^= 1 << i // устанавливаем i-й бит в 0, AND NOT обнуляет бит в числе; конъюнкция с отрицанием
	fmt.Printf("%d-й бит установлен в 0: %d\nДвоичное представление: %0b\n", i, number, number)
}

/*
	Дизъюнкция работает так, что если хоть один из операндов истинный,
	то значение тоже будет истинным. Конъюнкция с отрицанием всегда ложна,
	если второй операнд истинный.
*/

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа inter{}.
*/

package main

import (
	"fmt"
)

func getType(i interface{}) {
	switch i.(type) { // переключатель типов, который сравнивает типы
	case int:
		fmt.Println("value:", i, "type is int", "\n")
	case string:
		fmt.Println("value:", i, "type is string", "\n")
	case bool:
		fmt.Println("value:", i, "type is bool", "\n")
	case chan int:
		fmt.Println("value:", i, "type is channel", "\n")
	default:
		fmt.Println("value:", i, "unknown type", "\n")
	}
}

func main() {
	getType(1)
	getType("asd")
	getType(true)

	i := make(chan int)
	getType(i)

	j := []int{1, 2, 3}
	getType(j)
}

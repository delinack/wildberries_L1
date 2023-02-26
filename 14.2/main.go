/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа inter{}.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}

	i = 1
	fmt.Println("value:", i, "type:", reflect.TypeOf(i), "\n")

	i = "asd"
	fmt.Println("value:", i, "type:", reflect.TypeOf(i), "\n")

	i = true
	fmt.Println("value:", i, "type:", reflect.TypeOf(i), "\n")

	i = make(chan int)
	fmt.Println("value:", i, "type:", reflect.TypeOf(i), "\n")
	// функция reflect.TypeOf() принимает пустой интерфейс, а возвращает тип данного
}

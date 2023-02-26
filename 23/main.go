/*
Удалить i-ый элемент из слайса.
*/

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2 // удаляем третий элемент

	fmt.Println(slice)
	slice = append(slice[:index], slice[index+1:]...) // изменим срез, исключив элемент с индексом index, ... для раскрытия слайса
	fmt.Println(slice)
}

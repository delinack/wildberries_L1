/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{5, 2, 8, 4, 7, 1, 3, 9, 6}
	fmt.Println("before:", array)

	sort.Slice(array, func(i, j int) bool { // используем функцию сортировки слайсов из пакета sort
		return array[i] < array[j] // зададим параметр сортировки по возрастанию
	})

	fmt.Println("after 1:", array)

	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j] // теперь по убыванию
	})

	fmt.Println("after 2:", array)
}



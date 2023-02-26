/*
Реализовать бинарный поиск встроенными методами языка.
*/

package main

import "fmt"

func binarySearch(array []int, target int) int {
	left := 0               // первый индекс
	right := len(array) - 1 // последний индекс

	for left <= right { // пока остались не перебранные элементы
		mid := (left + right) / 2 // всегда начинаем поиск цели с середины
		if array[mid] == target { // если цель найдена, возвращаем ее
			return mid
		} else if array[mid] < target { // если цель правее центра, отбрасываем левую часть
			left = mid + 1
		} else {
			right = mid - 1 // если цель левее центра, отбрасываем правую часть
		}
		// в следующей итерации ищем в одной из двух половин новый центр
	}

	return -1 // возвращаем, если цель не найдена
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // бинарный поиск производится по отсортированному массиву
	target := 4
	fmt.Printf("index of target is: %d\n", binarySearch(array, target)) // передаем отсортированный массив
}

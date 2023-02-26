/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import "fmt"

func intersection(a, b []int) []int {
	m := make(map[int]bool)
	var res []int

	for _, val := range a {
		m[val] = false // добавим значения из первого массива в мапу, сделав их ключами
	}

	for _, val := range b {
		if _, ok := m[val]; ok { // если в мапе найдется такой ключ, то добавим значение в пересечение
			res = append(res, val)
		}
	}

	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(intersection(a, b))

	a1 := []int{1, 8, 5, 3, 5, 4, 4, 0}
	b1 := []int{5, 4, 4, 0}
	fmt.Println(intersection(a1, b1))
}

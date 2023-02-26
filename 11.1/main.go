/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	var res []int

	for i := range a { // запустим вложенный цикл
		for j := range b {
			if a[i] == b[j] { // если в множестве b есть такой же элемент, как и в a, - добавим его в пересечение
				res = append(res, a[i])
			}
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

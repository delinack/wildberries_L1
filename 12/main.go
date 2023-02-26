/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool)
	for _, str := range strings {
		m[str] = true // в мапу будут записаны только оригинальные ключи, что не допустит повторений
	}
	fmt.Println(m)
}

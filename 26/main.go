/*
Разработать программу, которая проверяет, что все символы в
строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true abCdefAaf — false aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	for _, i := range str {
		if strings.Count(str, string(i)) > 1 { // воспользуемся функцие подсчета символов в строке
			return false // если символов больше 1, то возвращаем false
		}
	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "bCdefA"
	fmt.Printf("%s – %t\n%s – %t\n%s – %t\n%s – %t\n",
		s1, unique(strings.ToLower(s1)), s2, unique(strings.ToLower(s2)), s3, unique(strings.ToLower(s3)), s4, unique(strings.ToLower(s4)))
	// перед отправкой строк в функцию, приведем их к нижнему регистру
}

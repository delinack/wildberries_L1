/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"strings"
)

func wordsReverser(str string) string {
	split := strings.Split(str, " ") // создадим массив из слов
	strs := make([]string, len(split))
	for i := 0; i < len(split); i++ {
		strs[i] = split[len(split)-i-1] // в новый массив будем записывать слова из старого в обратном порядке
	}
	return strings.Join(strs, " ") // соеденим слова обратно в строку
}

func main() {
	str := "snow dog sun"
	fmt.Println(wordsReverser(str))
}

// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

package main

import "fmt"

func reverseStr(str string) string {
	rn := []rune(str) // создадим руну
	str = ""
	for i := len(rn) - 1; i != -1; i-- { // бдем итерироваться по строке с конца
		str += string(rn[i]) // запишем в новую строку
	}
	return str
}

func main() {
	fmt.Println(reverseStr("главрыба"))
	fmt.Println(reverseStr("🦜 🐧 🦉"))
}

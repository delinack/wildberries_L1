/*
//Разработать программу, которая перемножает,
//делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
	_ "math/big"
)

func main() {
	a, b := big.NewInt(1), big.NewInt(1) // создадим переменные типа *Int

	a.Exp(big.NewInt(2), big.NewInt(45), nil) // зададим значение 2^45
	b.Exp(big.NewInt(2), big.NewInt(35), nil) // зададим значение 2^35

	fmt.Printf("a + b: %d,\na - b: %d,\na * b: %d,\na / b: %d.\n",
		new(big.Int).Add(a, b), new(big.Int).Sub(a, b), new(big.Int).Mul(a, b), new(big.Int).Div(a, b))
	// воспользуемся функциями сложения, вычитания, умножения и деления из пакета big для работы с большими числами
}

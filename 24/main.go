/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct { // структура с инкапсулированными полями
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y} // конструктор
}

func Distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x               // нахождение длин катетов
	dy := p1.y - p2.y               // нахождение длин катетов
	return math.Sqrt(dx*dx + dy*dy) // длина гипотенузы
}

func main() {
	p1 := NewPoint(1, 2) // создание экземпляров структуры Point
	p2 := NewPoint(4, 6)

	// поля p1 и p2 невозможно вызвать в этом пакете из-за инкапсуляции
	fmt.Printf("Расстояние между точками %v\n", Distance(p1, p2))
}

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от
родительской структуры Human (аналог наследования).
*/

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human // встраивание структуры (наследование)
	Act   string
}

func (a *Action) DoSomething() { // метод структуры Action
	fmt.Printf("Human %s, %d is %s.\n", a.Name, a.Age, a.Act) // вызов полей Action (в том числе и унаследованные поля Human)
}

func main() {
	human1 := Action{Human{Name: "Jack", Age: 27}, "sleeping"} // создаем экземпляр структуры, передаем значения
	human1.DoSomething()                                       // вызываем метод данного экземпляра

	human1.Name = "Paul" // изменим значения полей уже существующего экземпляра
	human1.Act = "eating"
	human1.DoSomething()

	human2 := Action{Human{Name: "Rose", Age: 23}, "learning"}
	human2.DoSomething()
}

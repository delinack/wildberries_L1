/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"fmt"
)

func main() {
	quit := make(chan int)
	work := make(chan int)
	go func() {
		for {
			select {
			case <-quit:
				return
			case data := <-work:
				fmt.Println("Working... ", data) // задаем дефолтное выполнение
			}
		}
	}()

	for i := 0; i < 10; i++ {
		work <- i + 1 // читаем в канал в цикле
	}
	quit <- 0 // по завершению цикла читаем в канал выхода
}

/*
	Мультиплексирование каналов с помощью select
	Горутина завершает работу, когда срабатывает case <-quit

	Завершение работы горутины при передачи в канал значения
*/

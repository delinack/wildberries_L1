/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	select {
	case <-time.After(duration): // канал time.After(duration), который отправит сигнал через указанный интервал времени
		return
	}
}

func main() {
	fmt.Println("start")

	sleep(time.Second)

	fmt.Println("1 second")

	sleep(500 * time.Millisecond)

	fmt.Println("500 milliseconds")

	sleep(3 * time.Second)

	fmt.Println("3 seconds")

}

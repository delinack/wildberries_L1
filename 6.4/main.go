/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println("Working...")
		runtime.Goexit() // немедленное завершение выполнения горутины
	}()
	time.Sleep(time.Millisecond)
}

/*
	runtime.Goexit() принудительно завершает выполнение текущей горутины.
*/

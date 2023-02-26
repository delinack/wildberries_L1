/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // контекст для отмены

	go func() {
		fmt.Printf("Working...\n")
		<-ctx.Done() // ожидаем сигнала контекста
	}()

	time.Sleep(time.Millisecond) // устанавливаем задержу, чтобы горутина успела напечатать сообщение
	cancel()                     // завершаем выполнение горутины
}

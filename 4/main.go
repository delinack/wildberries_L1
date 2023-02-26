/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные
данные из канала и выводят в stdout. Необходима возможность
выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func working(inChan chan int, closeChan chan bool, id int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик на 1
	for {
		select { // мультиплексиованние каналов
		case <-closeChan: // слушаем сигнал остановки на канале.
			return
		case data := <-inChan: // при срабатывании этого кейса происходит чтение данных из канала
			fmt.Printf("Worker %d : %d\n", id+1, data) // печать данных с указанием воркера
		}
		runtime.Gosched() // передача выполнения на каждой итерации цикла с одной горутины на другую для чередующегося вывода
	}
}

func main() {
	dataChan := make(chan int)   // канал для чтения произвольных данных
	closeChan := make(chan bool) // канал для закрытия и завершения цикла

	wg := &sync.WaitGroup{}                // создаем ссылку на wait group
	n := flag.Int("workers", 0, "workers") // храние значения, переданного флагу -workers

	flag.Parse() // парсинг флага
	for i := 0; i < *n; i++ {
		wg.Add(1)                              // добавляем воркер, кладем горутину в wait-группу, счетчик - 1
		go working(dataChan, closeChan, i, wg) // запуск горутины
	}

	go func() {
		for {
			dataChan <- rand.Int() // запись рандомных чисел в канал
		}
	}()

	c := make(chan os.Signal)                       // канал для отслеживания сигналов
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // отправляет сигнал о прерывании при нажатии Ctrl+C
	go func() {
		<-c
		close(closeChan) // закрываем канал, чтобы мультиплексор вышел из горутины
	}()

	wg.Wait() // ожидаем, пока счетчик не станет 0
}

/*
	Запускать с флагом -workers N, где N - количество воркеров.
*/

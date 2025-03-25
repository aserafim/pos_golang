package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	wait := sync.WaitGroup{}
	wait.Add(10)

	go publish(channel)
	go reader(channel, &wait)

	wait.Wait()

}

func publish(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}
	// Evita o deadlock
	close(ch)
}

func reader(ch chan int, w *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Recebido %d\n", x)
		w.Done()
	}
}

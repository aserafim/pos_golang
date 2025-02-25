package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	reader(channel)

}

func publish(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}
	// Evita o deadlock
	close(ch)
}

func reader(ch chan int) {
	for i := 1; i < 20; i++ {
		x := <-ch
		fmt.Printf("Recebido %d\n", x)
	}
}

package main

import (
	"fmt"
	"time"
	//"time"
)

func work_out(worker int, data <-chan int) {
	for i := range data {
		fmt.Printf("Worker %d processou o job %d\n", worker, i)
		//time.Sleep(time.Second)

	}
}

func main() {
	inicio := time.Now()
	data := make(chan int)
	//go work_out(1, data)

	for i := 0; i <= 1000; i++ {
		go work_out(i, data)
	}

	for i := 2; i <= 1000000; i++ {
		data <- i
	}

	fim := time.Now()
	fmt.Println(inicio)
	fmt.Println(fim)

}

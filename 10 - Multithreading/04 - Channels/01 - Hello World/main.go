package main

import (
	"fmt"
)

// Thread 1
func main() {

	ch1 := make(chan string)

	// Thread 2
	go func() {
		ch1 <- "Hello, World!"
		ch1 <- "Tentando encher o channel"
		fmt.Println("Isso não será impresso na tela...")		
	}()

	// Thread 1
	str := <-ch1
	fmt.Println(str)

}

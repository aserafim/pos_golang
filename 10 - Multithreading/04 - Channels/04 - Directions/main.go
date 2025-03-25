package main

import (
	"fmt"
)

// Função que só pode
// enviar dados para o canal
func send(nome string, data chan<- string) {
	data <- nome
}

// Função que só pode
// ler/receber dados
func receive(data <-chan string) {
	fmt.Println(<-data)
}

func main() {

	hello := make(chan string)
	go send("Hello", hello)
	receive(hello)

}

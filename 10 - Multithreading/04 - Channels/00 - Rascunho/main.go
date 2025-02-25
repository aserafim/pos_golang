package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	// Goroutine 1 tenta enviar dois valores
	go func() {
		fmt.Println("Enviando mensagem 1...")
		ch <- "Mensagem 1" // 🚨 OK, pois o `main()` receberá esse valor
		fmt.Println("Enviando mensagem 2...")
		ch <- "Mensagem 2" // 🚨 BLOQUEIA aqui! (ninguém está lendo)
		fmt.Println("Isso nunca será impresso!") // Nunca executa
	}()

	// Recebe apenas UM valor
	msg := <-ch
	fmt.Println("Recebido:", msg)
}

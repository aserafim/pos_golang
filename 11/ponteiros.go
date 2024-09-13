package main

import "fmt"

func main() {

	fmt.Println("Ponteiros")

	// Memória -> endereço -> Valor
	a := 10
	println(a)
	println(&a)

	// Sempre que virmos um * seguido
	// de um tipo estamos falando
	// de um endereço de memória
	var ponteiro *int = &a
	println(ponteiro)

	// Sempre que virmos um *
	// seguido de um nome de variável
	// estamos falando do que está
	// armazenado naquela variável
	*ponteiro = 20
	println(a)
}

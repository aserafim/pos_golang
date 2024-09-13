package main

import "fmt"

func addDez(a *int) {
	*a += 10
}

func main() {

	fmt.Println("Ponteiros")

	// Memória -> endereço -> Valor
	a := 10
	println(a)

	// Sempre que virmos um & seguido
	// de um nome de variável
	// estamos nos referindo ao
	// nome do endereço (código hash)
	println(&a)

	// Sempre que virmos um * seguido
	// de um tipo estamos falando
	// de um ponteiro para uma variável
	var ponteiro *int = &a
	println(ponteiro)

	// Sempre que virmos um *
	// seguido de um nome de variável
	// estamos falando do que está
	// armazenado naquela variável
	*ponteiro = 0
	println(a)

	addDez(&a)
	println(a)
}

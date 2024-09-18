package main

import "fmt"

func main() {
	fmt.Println("Loops")

	// Tipos de for mais usados

	// for clássico
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// similar ao while temos
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// utilizando o range
	nomes := []string{"Alefe", "Natália", "Alice", "Lulu", "Nick", "Alícia"}

	for k, v := range nomes {
		fmt.Println(k+1, v)
	}

	// loop infinito
	// uteis quando vamos esperar
	// retornos de requisições

	for {
		fmt.Println("Ao infinito e além...")
	}
}

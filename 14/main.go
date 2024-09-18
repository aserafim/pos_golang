package main

import "fmt"

func main() {
	fmt.Println("Loops")

	//Tipos de for mais usados

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	nomes := []string{"Alefe", "Natália", "Alice", "Lulu", "Nick", "Alícia"}

	for k, v := range nomes {
		fmt.Println(k, v)
	}
}

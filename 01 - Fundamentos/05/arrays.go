package main

import "fmt"

func main() {
	fmt.Println("Tipos compostos")

	var meuArray [3]int
	meuArray[0] = 2
	meuArray[1] = 4
	meuArray[2] = 8

	fmt.Println(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d.\n", i, v)
	}

}

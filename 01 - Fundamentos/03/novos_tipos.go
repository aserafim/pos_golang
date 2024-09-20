package main

import "fmt"

type ID int

func main() {
	var id ID = 1

	fmt.Println(id)

	fmt.Printf("Tipo %T\n", id)

}

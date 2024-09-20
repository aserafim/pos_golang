package main

import "fmt"

func main() {

	fmt.Println("Maps")

	//var meuMapa = map[int]int{}

	salarios := map[string]float32{
		"Alefe":   8000.34,
		"Natalia": 11000,
	}

	fmt.Println(salarios["Natalia"])

	//salarios2 := make(map[int]int)

	//var array [21]int
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %.2f.\n", nome, salario)
	}

	//array2 := [3]int{1,2,3}

}

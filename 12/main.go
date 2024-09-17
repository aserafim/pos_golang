package main

import "fmt"

// Função que soma os valores inteiros
// de uma estrutura map
func SomaInteiros(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

// E se quiséssemos somar também
// valores float32 ou float64??
func SomaValores[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {

	m := map[string]int{"Alefe": 500, "Natalia": 1000, "Antonieta": 4000}
	fmt.Println(SomaInteiros(m))

	m2 := map[string]float64{"Alefe": 500.9, "Natalia": 1000.8, "Antonieta": 4000.7}
	fmt.Printf("O valor final é %.2f\n", SomaValores(m2))

}

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

// Podemos ainda trabalhar com
// constraints
type Number interface {
	~int | float32 | float64
}

func SomaValoresConstraints[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Nem todo int é um int
// MyNumber é um int
// Porém não será aceito como
// um Number (constraint)
// a menos que coloquemos
// um ~ no como prefixo
// em nossa constraint
type MyNumber int

// Para compararmos Constraints
// podemos usar uma constraint
// específica, comparable
// desse modo o Go permite
// utilizar quaisquer valores
// como parâmetros desde que
// eles sejam valores comparáveis
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Alefe": 500, "Natalia": 1000, "Antonieta": 4000}
	fmt.Println(SomaInteiros(m))

	m2 := map[string]float64{"Alefe": 500.9, "Natalia": 1000.8, "Antonieta": 4000.7}
	fmt.Printf("O valor final é %.2f\n", SomaValores(m2))

	m3 := map[string]float32{"Alefe": 500.9, "Natalia": 1000.8, "Antonieta": 4000.7}
	fmt.Printf("O valor final é %.2f\n", SomaValoresConstraints(m3))

	//Esse trecho funciona porque foi colocado um ~ na definição da constraint
	m4 := map[string]MyNumber{"Alefe": 500, "Natalia": 1000, "Antonieta": 4000}
	fmt.Printf("O valor final é %d\n", SomaValoresConstraints(m4))

	//Funcionamento da função com comparable
	fmt.Println(Compara("a", "b"))
	fmt.Println(Compara(1, 1))
	fmt.Println(Compara(1.3, 1.5))
	fmt.Println(Compara(1, 1.0))
	//fmt.Println(Compara(1,"1.0")) Tipos não comparáveis -- não compila!

}

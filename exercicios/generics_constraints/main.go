/*
1. Função Genérica para Encontrar o Maior Elemento
Crie uma função genérica que recebe um slice de qualquer
tipo numérico e retorna o maior valor presente nele.
Utilize constraints para garantir que o tipo
genérico seja um número (inteiro ou float).
*/

/*
2. Implementação de uma Fila Genérica
Implemente uma fila (queue) usando uma
estrutura de dados genérica.
A fila deve ter as funções de enfileirar (enqueue),
desenfileirar (dequeue) e retornar o tamanho (size).
Teste a fila com diferentes tipos,
como int, string e float64.
*/

/*
3. Função Genérica de Comparação
Escreva uma função genérica que aceita
dois valores de um tipo genérico e retorna
o maior entre eles. Use constraints para
permitir apenas tipos que suportem a comparação.
*/

/*
4. Função Genérica de Soma
Crie uma função genérica que recebe
um slice de números (inteiros ou floats)
e retorna a soma de todos os elementos.
*/

/*
5. Mapa Genérico
Implemente uma função genérica
que recebe um mapa de chaves e valores
de tipos genéricos e retorna uma lista
de todas as chaves.
*/

/*
6. Ordenação de Slice Genérico
Crie uma função genérica para ordenar
slices de qualquer tipo que suporte
operações de comparação, como int,
float64 ou string. Utilize constraints
adequadas para restringir os tipos permitidos.
*/

package main

import (
	"fmt"
)

// Exercício 1
type Number interface {
	~int | ~float64 | ~float32
}

func partition[T Number](arr []T, low, high int) ([]T, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort[T Number](arr []T, low, high int) ([]T, T) {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr, _ = quickSort(arr, low, p-1)
		arr, _ = quickSort(arr, p+1, high)
	}
	return arr, arr[len(arr)-1]
}

// Exercício 3
func RetornaMaior[T Number](num1 T, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}

// Exercício 4
func Soma[T Number](valores []T) T {
	var soma T
	for _, num := range valores {
		soma += num
	}

	return soma
}

func main() {
	fmt.Println("Exercícios")

	//Testes Exercício 1
	array := []int{50, 40, 30, 20, 10}
	arrayDois := []int{5.0, 4.0, 3.0, 2.0, 1.0}

	_, maior := quickSort(array, 0, len(array)-1)
	fmt.Println(maior)

	_, maiorDois := quickSort(arrayDois, 0, len(arrayDois)-1)
	fmt.Println(maiorDois)

	//Testes Exercício 3
	fmt.Println(RetornaMaior(1.0, 8))

	//Testes Exercício 4
	fmt.Println(Soma(array))
	fmt.Println(Soma(arrayDois))

}

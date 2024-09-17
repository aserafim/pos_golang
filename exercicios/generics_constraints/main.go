package main

import (
	"fmt"
)

/*
1. Função Genérica para Encontrar o Maior Elemento
Crie uma função genérica que recebe um slice de qualquer
tipo numérico e retorna o maior valor presente nele.
Utilize constraints para garantir que o tipo
genérico seja um número (inteiro ou float).
*/

/*
3. Função Genérica de Comparação
Escreva uma função genérica que aceita dois valores de um tipo genérico e retorna o maior entre eles. Use constraints para permitir apenas tipos que suportem a comparação.

4. Função Genérica de Soma
Crie uma função genérica que recebe um slice de números (inteiros ou floats) e retorna a soma de todos os elementos.
*/

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

func main() {
	fmt.Println("Exercícios")

	array := []int{50, 40, 30, 20, 10}
	_, maior := quickSort(array, 0, len(array)-1)
	fmt.Println(maior)
}

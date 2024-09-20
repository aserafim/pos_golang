package main

import (
	"errors"
	"fmt"
	"strings"
)

// Exercício 1 - Função de Soma: Escreva uma função chamada sum que receba
// dois inteiros como parâmetros e retorne a soma deles.
func sum(a int, b int) int {
	return a + b
}

//Alternativa
// func sum(a, b int) int {
// 	return a + b
// }

// Exercício 5 - Função com Múltiplos Retornos: Crie uma função chamada
// divide que receba dois números inteiros e retorne o quociente e o resto
// da divisão entre eles.
func divide(a int, b int) (int, int) {
	if b != 0 {
		return (a / b), (a % b)
	}
	return 0, 0
}

// Funções com múltiplos retornos
// são úteis no tratamento de erros
func divWithErros(a int, b int) (int, error) {
	if b != 0 {
		return a / b, errors.New("Impossível dividir por zero")
	}
	return 0, nil
}

// Exercício 2 -Função de Verificação de Paridade: Crie uma função chamada
// isEven que receba um número inteiro e retorne true se o número for par,
// e false caso contrário.
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// Exercício 3 - Função de Fatorial: Implemente uma função chamada
// factorial que receba um número inteiro positivo e retorne o fatorial
// desse número.
func factorial(num int) int {
	result := 1
	if num == 1 {
		return 1
	} else if num < 1 {
		return 0
	}
	for i := num; i >= 1; i-- {
		result *= i
	}
	return result
}

// Exercício 4 - Função para Reverter String: Escreva uma função chamada
// reverseString que receba uma string como parâmetro e retorne a string
// invertida.
func reverseString(texto string) string {
	reversedString := make([]string, len(texto))
	println(len(texto))
	j := len(texto) - 1

	for i := 0; i < len(texto); i++ {
		reversedString[i] = string(rune(texto[j]))
		j--
	}
	return strings.Join(reversedString, "")
}

// Exercício 6 - Função Recursiva de Fibonacci: Implemente uma função
// chamada fibonacci que receba um número inteiro n e retorne o enésimo
// número da sequência de Fibonacci de forma recursiva.

// Exercício 7 - Função para Ordenar uma Lista: Escreva uma função chamada
// sortArray que receba um slice de inteiros e retorne esse slice ordenado.

// Exercício 8 - Função para Encontrar o Maior Número: Crie uma função
// chamada maxNumber que receba um slice de inteiros e retorne o maior
// número presente no slice.

// Funções variádias são úteis quando não sabemos
// quantos itens teremos nos parâmetros
// das nossas funções
func variadica(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {

	fmt.Println("Funções")

	// println(sum(1, 2))

	// println(div(2, 0))

	// println(divWithErros(2, 0))

	//println(isEven(0))

	//println(factorial(1))

	//println(reverseString("alefe"))

	println(divide(7, 2))

}

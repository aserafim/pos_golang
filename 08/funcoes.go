package main

import (
	"errors"
	"fmt"
)

//Exercício 1 - Função de Soma: Escreva uma função chamada sum que receba 
//dois inteiros como parâmetros e retorne a soma deles.
func sum(a int, b int) int {
	return a + b
}

//Alternativa
// func sum(a, b int) int {
// 	return a + b
// }

// Funções com múltiplos retornos
// são úteis no tratamento de erros
func div(a int, b int) (int, bool) {
	if b != 0 {
		return a / b, true
	}
	return 0, false
}

// Funções com múltiplos retornos
// são úteis no tratamento de erros
func divWithErros(a int, b int) (int, error) {
	if b != 0 {
		return a / b, errors.New("Impossível dividir por zero")
	}
	return 0, nil
}

//Exercício 2 -Função de Verificação de Paridade: Crie uma função chamada 
//isEven que receba um número inteiro e retorne true se o número for par, 
//e false caso contrário.



func main() {

	fmt.Println("Funções")

	println(sum(1, 2))

	println(div(2, 0))

	println(divWithErros(2, 0))
}

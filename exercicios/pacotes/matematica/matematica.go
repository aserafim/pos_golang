/*
1. Criando um pacote simples:
  - Crie um pacote chamado `matematica` com funções para somar, subtrair, multiplicar e dividir dois números.
  - No pacote principal, importe e utilize essas funções.
*/

package matematica

type Number interface {
	~int | ~float32 | ~float64
}

func Soma[T Number](num1 T, num2 T) T {
	return num1 + num2
}

func Divisao[T Number](num1 T, num2 T) T {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}

func Subtracao[T Number](num1 T, num2 T) T {
	return num1 - num2
}

func Multiplicacao[T Number](num1 T, num2 T) T {
	return num1 * num2
}

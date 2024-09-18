/*
Exercícios de Pacotes em GoLang:

1. Criando um pacote simples:
  - Crie um pacote chamado `matematica` com funções para somar, subtrair, multiplicar e dividir dois números.
  - No pacote principal, importe e utilize essas funções.

2. Estrutura de múltiplos pacotes:
  - Crie dois pacotes: `geometria` e `trigonometria`.
  - Em `geometria`, implemente funções para calcular área e perímetro de figuras geométricas como círculo, quadrado e retângulo.
  - Em `trigonometria`, crie funções que calculam seno, cosseno e tangente de um ângulo.
  - No pacote principal, use funções de ambos os pacotes para realizar cálculos.

3. Gerenciamento de dependências:
  - Crie um projeto que dependa de um pacote externo (pode ser um pacote da biblioteca padrão ou um pacote de terceiros).
  - Utilize o Go Modules para gerenciar as dependências.

4. Agrupando funcionalidades:
  - Crie um pacote chamado `conversores` com funções que convertem unidades (ex: metros para quilômetros, gramas para quilogramas, Celsius para Fahrenheit).
  - No pacote principal, importe e utilize essas funções.

5. Pacote utilitário:
  - Crie um pacote utilitário chamado `stringsutil` que tenha funções para manipulação de strings, como inverter uma string, contar vogais e verificar se uma string é palíndroma.
  - Escreva testes para essas funções no pacote `stringsutil_test`.
*/
package main

import (
	"ex1/matematica"
	"fmt"
)

func main() {
	fmt.Println("Exercícios Pacotes")

	//Testando Ex 1
	fmt.Println(matematica.Soma(1, 2))

	fmt.Println(matematica.Divisao(4, 2))

	fmt.Println(matematica.Multiplicacao(2, 8))

	fmt.Println(matematica.Subtracao(4, 9))
}

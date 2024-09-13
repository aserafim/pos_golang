/*
Ponteiros Simples: Crie uma função chamada increment que receba um ponteiro para um inteiro e incremente o valor desse inteiro. Teste essa função em um código onde a variável original é modificada corretamente usando o ponteiro.
*/

/*
Ponteiros e Structs: Crie uma struct chamada Person com os campos name e age. Escreva uma função que receba um ponteiro para a struct Person e altere o valor da idade (age) para 30. Verifique se a struct original é atualizada corretamente.
*/

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Trintou() {
	p.age = 30
}

func increment(inteiro *int) {
	*inteiro += 1
}

func main() {
	fmt.Println("Ponteiros Simples")

	num := 10
	fmt.Printf("Valor antes do incremento: %d.\n", num)

	increment(&num)
	fmt.Printf("Valor após o incremento: %d.\n", num)

	p1 := Person{
		name: "Alefe",
		age:  29,
	}
	fmt.Printf("A idade atual é %d.\n", p1.age)
	p1.Trintou()
	fmt.Printf("A nova idade é %d.\n", p1.age)

}

/*
Ponteiros Simples: Crie uma função chamada increment que receba um ponteiro para um inteiro e incremente o valor desse inteiro. Teste essa função em um código onde a variável original é modificada corretamente usando o ponteiro.
*/

/*
Ponteiros e Structs: Crie uma struct chamada Person com os campos name e age. Escreva uma função que receba um ponteiro para a struct Person e altere o valor da idade (age) para 30. Verifique se a struct original é atualizada corretamente.
*/

/*
Combinação de Interfaces e Ponteiros: Crie uma interface Animal com um método Speak() string. Implemente essa interface para diferentes animais, como Dog, Cat, etc. Em seguida, crie um método que receba um ponteiro para uma dessas structs e modifique algum comportamento interno, como o nome do animal, e veja como isso afeta o método Speak().
*/

package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct {
	nome  string
	idade int
}

func (c Cat) Speak() {
	fmt.Println("Miau!!")
}

type Dog struct {
	nome  string
	idade int
}

func (d Dog) Speak() {
	fmt.Println("Au! Au!")
}

func (d *Dog) AlteraNome(novo_nome string) {
	d.nome = novo_nome
}

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

	c1 := Cat{nome: "Severino", idade: 2}
	d1 := Dog{nome: "Sami", idade: 4}

	fmt.Println(d1)
	d1.Speak()

	fmt.Println(c1)
	c1.Speak()

	d1.AlteraNome("Pepeu")
	fmt.Println(d1)
	d1.Speak()

}

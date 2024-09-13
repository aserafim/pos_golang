/*
Interfaces e Polimorfismo: Crie uma interface chamada Shape que tenha um método Area() float64. Depois, crie duas structs: Circle e Rectangle. Ambas devem implementar a interface Shape. O Circle terá um raio (radius) e o Rectangle terá largura (width) e altura (height). Escreva um código que calcule a área de diferentes formas usando a interface Shape.
*/

/*
Interface com Métodos e Ponteiros: Crie uma interface chamada Database com os métodos Connect() e Disconnect(). Implemente essa interface para diferentes tipos de bancos de dados, como SQLDatabase e NoSQLDatabase, usando structs. Faça com que a função Connect altere um valor dentro da struct através de um ponteiro, representando o status da conexão.
*/

package main

import (
	"fmt"
	"math"
)

type Cat struct {
	nome  string
	idade int
}

func (c Cat) Speak() {
	fmt.Println("Miau!!")
}

type Database interface {
	Connect()
	Disconnect()
}

type SQLDatabase struct {
	Name      string
	Conectado bool
}

func (banco *SQLDatabase) Connect() {
	banco.Conectado = true

}

func (banco *SQLDatabase) Disconnect() {
	banco.Conectado = false

}

type NoSQLDatabase struct {
	Name      string
	Conectado bool
}

func (banco *NoSQLDatabase) Connect() {
	banco.Conectado = true

}

func (banco *NoSQLDatabase) Disconnect() {
	banco.Conectado = false

}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func main() {

	fmt.Println("Alguns exercícios")

	circuloUm := Circle{
		radius: 12.8,
	}

	retanguloUm := Rectangle{
		height: 4,
		width:  4,
	}

	fmt.Printf("A área do círculo é %.2f\n", circuloUm.Area())
	fmt.Printf("A área do retângulo é %.2f\n", retanguloUm.Area())

	bancoUm := NoSQLDatabase{
		Name:      "MongoDB",
		Conectado: false,
	}

	fmt.Println(bancoUm)
	bancoUm.Connect()
	fmt.Println(bancoUm)
	bancoUm.Disconnect()
	fmt.Println(bancoUm)

	bancoDois := SQLDatabase{
		Name:      "MySQL",
		Conectado: false,
	}

	fmt.Println(bancoDois)
	bancoDois.Connect()
	fmt.Println(bancoDois)
	bancoDois.Disconnect()
	fmt.Println(bancoDois)
}

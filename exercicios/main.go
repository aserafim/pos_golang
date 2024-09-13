package main

import (
	"fmt"
	"math"
)

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

}

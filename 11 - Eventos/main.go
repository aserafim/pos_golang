package main

import "fmt"

type Pessoa struct {
	Name       string
	Idade      int
	Peso       float32
	Experienca map[string]int
}

func NewPessoa() *Pessoa {
	return &Pessoa{
		Name:       "null",
		Idade:      0,
		Peso:       0.0,
		Experienca: make(map[string]int),
	}
}

func main() {
	mapa := make(map[string]int)

	//append(mapa[alefe]1)

	mapa["alefe"] = 1
	mapa["natalia"] = 2
	mapa["olabone"] = 3
	mapa["vizinho"] = 4

	_, ok := mapa["natalia"]
	if ok {
		fmt.Println("Encontrado")
	}

	alefe := NewPessoa()
	fmt.Println(alefe.Name)
}

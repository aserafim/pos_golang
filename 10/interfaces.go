package main

import "fmt"

// Note que uma interface aceita
// apenas MÉTODOS em sua assinatura
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.", c.Nome)
}

type Aluno struct {
	nome  string
	notas []float64
	ativo bool
}

func (a Aluno) Desativar() {

}

func (aluno Aluno) calculaMedia() float64 {
	var media float64
	for _, nota := range aluno.notas {
		media += nota
	}

	return media / float64(len(aluno.notas))
}

func main() {

	fmt.Println("Interfaces")

	clienteUm := Cliente{
		Nome:  "Alefe",
		Idade: 30,
		Ativo: true,
	}

	fmt.Println(clienteUm)

	clienteUm.Desativar()
	fmt.Println(clienteUm)

	alunoUm := Aluno{
		nome:  "Alfredo",
		notas: []float64{1, 3, 4},
		ativo: true,
	}

	alunoUm.Desativar()

}

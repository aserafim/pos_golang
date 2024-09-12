package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

type Carro struct {
	marca  string
	modelo string
	ano    int
}

func compara(carroUm Carro, carroDois Carro) bool {
	if carroUm.marca == carroDois.marca && carroUm.modelo == carroDois.modelo && carroUm.ano == carroDois.ano {
		return true
	}
	return false
}

type Livro struct {
	titulo  string
	autor   string
	ano     int
	qtd_pgs int
}

func (livro Livro) imprimir() {

	fmt.Println("Título: ", livro.titulo, ", escrito por ", livro.autor, "\nPublicado em ", livro.ano, ", ", livro.qtd_pgs, " páginas.")
}

type Produto struct {
	nome  string
	preco float32
}

func alteraPreco(p *Produto, novoValor float32) {
	p.preco = novoValor
}

type Aluno struct{
	nome string
	notas []float64
}


func main() {

	fmt.Println("Estruturas - Go Way of programming")

	carroUm := Carro{
		marca:  "Chevrolet",
		modelo: "Chevette",
		ano:    1965,
	}

	carroDois := Carro{
		marca:  "Fiat",
		modelo: "Uno",
		ano:    1995,
	}

	fmt.Println(carroUm.modelo, " - ", carroUm.marca)
	fmt.Println(carroDois.modelo, " - ", carroDois.marca)

	livroUm := Livro{
		titulo:  "1984",
		autor:   "George Orwell",
		ano:     1949,
		qtd_pgs: 416,
	}

	livroUm.imprimir()

	fmt.Println(compara(carroUm, carroDois))

	produtoUm := Produto{
		nome:  "Lápis",
		preco: 1.50,
	}

	fmt.Println("Produto: ", produtoUm.nome, " valor: R$", produtoUm.preco)

	alteraPreco(&produtoUm, 1.75)

	fmt.Println("Produto: ", produtoUm.nome, " valor: R$", produtoUm.preco)

	alunoUm = Aluno{
		nome: "Alefe",
		notas: [0]

	}

}

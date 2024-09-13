package main

import "fmt"

func addDez(a *int) {
	*a += 10
}

type Conta struct {
	saldo int
}

func (c Conta) simularEmprestimo(valor int) {
	c.saldo += valor
	fmt.Printf("Após a simulação o saldo será de %d.\n", c.saldo)
}

func (c *Conta) realizaEmprestimo(valor int) int {
	*&c.saldo += valor
	fmt.Printf("Emprestimo no valor de %d realizado com sucesso.\n", valor)
	return c.saldo
}

func main() {
	fmt.Println("Ponteiros")

	// Memória -> endereço -> Valor
	a := 10
	println(a)

	// Sempre que virmos um & seguido
	// de um nome de variável
	// estamos nos referindo ao
	// nome do endereço (código hash)
	println(&a)

	// Sempre que virmos um * seguido
	// de um tipo estamos falando
	// de um ponteiro para uma variável
	var ponteiro *int = &a
	println(ponteiro)

	// Sempre que virmos um *
	// seguido de um nome de variável
	// estamos falando do que está
	// armazenado naquela variável
	*ponteiro = 0
	println(a)

	addDez(&a)
	println(a)

	contaUm := Conta{
		saldo: 1000,
	}

	fmt.Println(contaUm.saldo)

	contaUm.simularEmprestimo(1200)
	fmt.Printf("O saldo real atual da contá é de %d.\n", contaUm.saldo)

	contaUm.realizaEmprestimo(3000)
	fmt.Printf("Saldo atual da conta: %d.\n", contaUm.saldo)
}

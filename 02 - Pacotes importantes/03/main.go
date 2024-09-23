package main

import (
	"encoding/json"
	"fmt"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {

	fmt.Println("Aula sobre json")

	conta := Conta{Numero: 1, Saldo: 4000}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))
}

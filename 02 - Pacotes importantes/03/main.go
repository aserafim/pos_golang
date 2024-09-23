package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	// Podemos ainda usar um encoder
	// que vai transformar em um JSON
	// e salvar num arquivo de saída
	// isso é útil quando queremos
	// realizar automaticamente a
	// "entrega" para algum servidor
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	// Agora vamos fazer o trabalho oposto
	// a desserialização
	// Note que aqui precisamos de uma
	// struct onde o json será armazenado
	// após o resultado da desserialização
	// ser finalizado
	jsonPuro := []byte(`{"Numero":3, "Saldo":8115}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}

	println(contaX.Numero)
	println(contaX.Saldo)

}

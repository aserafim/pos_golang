package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func getFromViaCep(ch chan int) {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição, %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s Localidade: %s Logradouro: %s\n", data.Cep, data.Localidade, data.Logradouro))

		fmt.Println("Arquivo gerado com sucesso.")
		fmt.Printf("CEP: %s Localidade: %s Logradouro: %s\n", data.Cep, data.Localidade, data.Logradouro)

	}
	time.Sleep(time.Second * 10)
	ch <- 1
}

func Correios(ch chan int) {
	ch <- 1
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	//var msg1 int

	go getFromViaCep(ch1)
	go Correios(ch2)

	select {

	case msg1 := <-ch1:
		fmt.Println("Recebido de ViaCep", msg1)

	case msg1 := <-ch2:
		fmt.Println("Recebido de Correios", msg1)
	}

}

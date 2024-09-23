package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Advice struct {
	Slip struct {
		ID     int    `json:"id"`
		Advice string `json:"advice"`
	} `json:"slip"`
}

func main() {

	req, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao realizar a requisição: %v\n", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
	}

	var advice Advice
	err = json.Unmarshal(res, &advice)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	fmt.Println(advice.Slip.Advice)

}

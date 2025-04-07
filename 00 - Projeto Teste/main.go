package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Advice struct {
	Code  string `json:"code"`
	Hello string `json:"hello"`
}

func main() {

	req, err := http.Get("https://hellosalut.stefanbohacek.dev/?ip=89.120.120.120")
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
	fmt.Println(advice.Hello)

}

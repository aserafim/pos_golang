package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Aula sobre chamadas HTTP")

	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

}

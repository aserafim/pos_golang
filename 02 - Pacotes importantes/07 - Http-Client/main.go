package main

import (
	"io"
	"net/http"
)

/*
  Uma forma de aumentar a performance
  de nossos sistemas é o definir
  o timeout das requisições
*/

func main() {
	//c := http.Client{Timeout: time.Microsecond}
	c:= http.Client{}
	resp, err := c.Post("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

/*
  Nesse caso o nosso panic retorna
  panic: Get "http://google.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
*/

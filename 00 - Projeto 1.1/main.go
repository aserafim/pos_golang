package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCep)
	http.ListenAndServe(":8080", nil)
}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}

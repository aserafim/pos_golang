package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	number := 1
	w.Write([]byte(fmt.Sprintf("You are the visitant of number %s", string(number))))
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":9000", nil)

}

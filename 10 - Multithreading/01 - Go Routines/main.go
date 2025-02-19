package main

import (
	"fmt"
	"time"
)

func callApi(apiName string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Calling %s", i+1, apiName+"\n")
	}
}

func main() {

	go callApi("Serasa")
	go callApi("FBI")
	time.Sleep(15 * time.Second)
}

package main

import (
	"fmt"
	"sync"
)

func callApi(apiName string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Calling %s", i+1, apiName+"\n")
		waitGroup.Done()
	}
}

func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(20)
	go callApi("Serasa", &waitGroup)
	go callApi("FBI", &waitGroup)
	waitGroup.Wait()
}

// O que são waitgroups?

// Como utiliamos eles?

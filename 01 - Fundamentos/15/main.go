package main

import "fmt"

func main() {

	fmt.Println("Condicionais")

	a := 0
	b := 1
	c := 2

	if a > b {
		fmt.Println("Maior")
	}

	if a > b && b > c {
		fmt.Println("Maior")
	}

	if a > b || b > c {
		fmt.Println("Maior")
	}

	switch a {
	case 0:
		fmt.Println(a)
	case 2:
		fmt.Println(b)
	default:
		fmt.Println("default")
	}
}

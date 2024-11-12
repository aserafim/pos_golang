package main

import (
	"fmt"

	"packaging/math"
)

func main() {
	fmt.Println("Working with Packaging")
	m := math.Math{A: 2, B: 4}
	fmt.Println(m.Add())
}

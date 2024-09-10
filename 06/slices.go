package main

import "fmt"

func main() {

	fmt.Println("Slices")

	slc := []int{10, 20, 30, 40, 50}

	fmt.Printf("len=%d, cap=%d %v\n", len(slc), cap(slc), slc)
}

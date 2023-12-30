package main

import "fmt"

func main() {
	fmt.Println("Pointer Study")

	// var ptr *int
	// fmt.Println("ptr: ", ptr)

	myNumber := 22

	var ptr = &myNumber
	fmt.Println("ptr: ", ptr)
	fmt.Println("ptr: ", *ptr)
	*ptr = *ptr + 2
	fmt.Println("new ptr: ", *ptr)
}

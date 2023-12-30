package main

import "fmt"

func main() {
	defer fmt.Println("Hello World") //A defer statement defers the execution of a function until the surrounding function returns. i.e., last in first out
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Defer Study")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

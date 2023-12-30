package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err ok
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error occur")
	}
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the rating is %T, ", input)
}

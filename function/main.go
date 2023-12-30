package main

import "fmt"

func greeter() {
	fmt.Println("Hello All!")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "is the sum of the number."
}

func main() {
	fmt.Println("Function in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult, myMessage := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(proResult, myMessage)

}

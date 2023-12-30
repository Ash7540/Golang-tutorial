package main

import "fmt"

func main() {
	fmt.Println("If-ELse Study")
	loginCount := 5
	var result string

	if loginCount < 5 {
		result = "Regular User"
	} else if loginCount > 5 {
		result = "Super User"
	} else {
		result = "Value is 5"
	}
	fmt.Println(result)


	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}
}

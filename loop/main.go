package main

import "fmt"

func main() {
	fmt.Println("Loop Study")

	// days := []string{"Sunday", "Tuesday", "Thursday", "Friday"}

	// fmt.Println(days)

	// for loop type 1
	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for loop type 2
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for loop type 3
	// for index, day := range days{
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }

	rougueValue := 1
	// for rougueValue < 10 {
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// 	// ++rougueValue will give error
	// }
	for rougueValue < 10 {
		if rougueValue == 2 {
			// break
			goto lvo
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++

	}

lvo:
	fmt.Println("Jumping at Ashu")

}

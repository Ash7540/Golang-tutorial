package main

import "fmt"

func main() {
	fmt.Println("Array Study")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Orange"

	// Accessing Array Elements
	fmt.Println("Fruit List: ", fruitList)
	fmt.Println("Fruit List Length is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Ladyfinger", "Beans"}
	fmt.Println("Veg List: ", vegList)
	fmt.Println("Veg List Length is: ", len(vegList))

}
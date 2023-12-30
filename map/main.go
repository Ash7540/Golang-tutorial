package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map study")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	// Accessing values in the map using keys
	fmt.Println(languages)
	fmt.Println("Js short for ",languages["JS"])
	
	// deleting the value 
	delete(languages, "JS")
	fmt.Println(languages)

	// loops in map
	for key, value := range languages{
		fmt.Printf("for key %v, value is %v\n", key, value )
	}

}

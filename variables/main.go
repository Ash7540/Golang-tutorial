package main

import (
	"fmt"
)


// jwtToken := 3000 not working as it is only allowed inside the method

const LoginToken string = "ashuchavan"  // Public



func main() {

	var username string = "Ashu"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool =  false
	fmt.Println(isLoggedIn )
	fmt.Printf("variable is of type: %T \n", isLoggedIn )

	var age uint8 =  22
	fmt.Println(age)
	fmt.Printf("variable is of type: %T \n", age )

	var marks float32 =  8.76
	fmt.Println(marks)
	fmt.Printf("variable is of type: %T \n", marks )

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)

	//implicit type

	var website = "github.com"
	fmt.Println(website)

	// no var style
	thirdVariable := 23.6
	fmt.Println(thirdVariable)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken )


}

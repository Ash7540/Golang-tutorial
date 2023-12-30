package main

import "fmt"


type User struct{
	Name string
	Email string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Struct Study")
	// no inheritance in golang; no super or parent

	ashu := User{"Ashu", "ashu123@gmail.com", true, 22}
	fmt.Println(ashu)
	fmt.Printf("Ashu details are:  %+v\n", ashu)
	fmt.Printf("Name is %v and email is %v.", ashu.Name, ashu.Email)

}
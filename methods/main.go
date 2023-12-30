package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active? ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.com"
	fmt.Println("Email of this user is: ", u.Email)
}

func main() {
	fmt.Println("Struct Study")
	// no inheritance in golang; no super or parent

	ashu := User{"Ashu", "ashu123@gmail.com", true, 22}
	// fmt.Println(ashu)
	// fmt.Printf("Ashu details are:  %+v\n", ashu)
	// fmt.Printf("Name is %v and email is %v.", ashu.Name, ashu.Email)
	ashu.GetStatus()
	ashu.NewMail()
	fmt.Printf("Name is %v and email is %v.", ashu.Name, ashu.Email)

}

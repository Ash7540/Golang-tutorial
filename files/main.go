package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files")
	content := "My name is Ashwin Chavan and I am invincible"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)


	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Txt data byte inside the file is \n", databyte)
	fmt.Println("Txt data inside the file is \n", string(databyte))
}


func checkNilErr(err error){
	if err != nil {
		panic(err)
	}

}
package main

import (
	"fmt"
	"net/url"
)


const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=uvhg6778fgheg"

func main() {
	fmt.Println("Handling URL in Golang")
	fmt.Println(myurl)

	//parsing the url
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)


	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n", qparams)


	fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	// for _, val := range qparams{
	// 	fmt.Println("Param is: ", val)
	// }

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath:  "user=ashu",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)


}

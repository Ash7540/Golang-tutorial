package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	fmt.Println("Content Length", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename" : "golang",
		"price" : "200"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)

	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// fake form payload
	data := url.Values{}
	data.Add("firstname", "Ashu")
	data.Add("lastname", "Chavan")
	data.Add("email", "ash123@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func main() {
	fmt.Println("Web Verb in Golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

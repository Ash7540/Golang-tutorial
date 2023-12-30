package main

import (
	"fmt"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB for API In Golang")
	r := router.Router()
	fmt.Println("Server is getting started ...")
	http.ListenAndServe(":8000", r)
	fmt.Println("Listening at port 8000...")
}
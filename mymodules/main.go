package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func greeter() {
	fmt.Println("Hello World")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

func main() {
	fmt.Println("My Module in Golang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	greeter()

	log.Fatal(http.ListenAndServe(":8000", r))
}

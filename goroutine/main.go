package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}

var wg sync.WaitGroup // usually use here pointer
var mut sync.Mutex    // usually use here pointer

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

func main() {
	fmt.Println("Concurrency and Goroutines in Golang")
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)

	}
	wg.Wait()
	fmt.Println(signals)

}

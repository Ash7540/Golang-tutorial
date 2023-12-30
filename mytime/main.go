package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello User")

	presentTime := time.Now()
	fmt.Println("The time is ", presentTime)
	fmt.Println("The time is ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.February, 18, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created Date is ", createdDate)
	fmt.Println("Created Date is ", createdDate.Format("01-02-2006 15:04:05 Monday"))

	
}
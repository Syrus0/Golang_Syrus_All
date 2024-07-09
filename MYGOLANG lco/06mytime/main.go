package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GOlang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // yo date chai format ma use garnai parcha
	createdDate := time.Date(2020, time.May, 23, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go ROutines")

	go greeter("Hello")
	greeter("world")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(5 * time.Microsecond)
		fmt.Println(s)
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int

	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber //referencign & is used for referencing
	fmt.Println("Value of Actual Pointer", ptr)
	fmt.Println("Value of Actual Pointer", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value", myNumber)

}

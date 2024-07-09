package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Funtions in Golang")
	greeterTwo()

	proRes, mess := proAdder(2, 5, 8, 7)

	fmt.Println("Pro REsult is ", proRes)
	fmt.Println("Pro message is ", mess)
	fmt.Println("Result is:", result)
}
func result(valOne, valTwo int) int {

	return valOne + valTwo

}

func proAdder(value ...int) (int, string) {
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "Jay Nepal"
}
func greeterTwo() {
	fmt.Println("Another Method")
}
func greeter() {
	fmt.Println("Namaste")
}

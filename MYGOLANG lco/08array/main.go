package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}

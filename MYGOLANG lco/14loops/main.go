package main

import "fmt"

func main() {

	fmt.Println("Loops in GoLang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])

	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// // }
	// for index, days := range days {
	// 	fmt.Printf("Index is %v and Value is %v", index, days)
	// }
	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto sy

		}

		if rogueValue == 5 {
			rogueValue++
			continue

		}
		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

sy:
	fmt.Println("Jumping to Syrus")
}

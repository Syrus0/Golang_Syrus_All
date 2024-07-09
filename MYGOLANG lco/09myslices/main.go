package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Litch", "Tomato")

	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3]) //slice ko part banaucha
	fmt.Println(fruitList)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	//highscores[4] = 777

	highscores = append(highscores, 555, 666, 321)

	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted((highscores))) // sort bhako cha ki chaina check garcha
	sort.Ints(highscores)                         //sorts in ascending order
	fmt.Println(highscores)
	// fmt.Println(sort.IntsAreSorted((highscores)))

	// remove a valur from slices based on index

	var courses = []string{"reactjs", "python", "swift", "ruby", "java", "GO"}
	fmt.Println(courses)

	var index int = 2                                       // 2 goes away
	courses = append(courses[:index], courses[index+1:]...) //index jancha'
	fmt.Println(courses)
}

// package main

// import "fmt"

// func main() {

// 	Syrus := []string{"Rajendra", "Bandana", "Sumi"}
// 	fmt.Printf("The type of value is %T\n", Syrus)
// 	fmt.Println(Syrus)

// 	Syrus = append(Syrus, "Chunte", "Sujal")
// 	fmt.Println(Syrus)

// 	Syrus = append(Syrus[:3])
// 	fmt.Println(Syrus)

// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Maps in Golang")

// 	languages := make(map[string]string)

// 	languages["JS"] = "Javascript"
// 	languages["RB"] = "Ruby"
// 	languages["PY"] = "Python"
// 	fmt.Println("list of all Languages: ", languages)
// 	fmt.Println("JS shorts for ", languages["JS"])

// 	delete(languages, "RB")
// 	fmt.Println("Languages", languages)

// 	//Looops in laps and they are interesting

// 	//for key, value := range languages {
// 	for _, value := range languages {
// 		//fmt.Printf("For key %v, value is %v \n", key, value)
// 		fmt.Printf("For key , value is %v \n", value)
// 	}

// }

package main

import "fmt"

func main() {

	fmt.Println("Learning Maps myself")

	Languages := make(map[string]string)

	Languages["G"] = "Nigga"
	Languages["Sy"] = "Syrus"

	fmt.Println(Languages)
	for key, value := range Languages {
		fmt.Printf("The short form for %v is %v\n", key, value)
	}

}

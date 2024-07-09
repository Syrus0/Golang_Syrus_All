package main

import "fmt"

func main() {

	fmt.Println("Structs in goLang")

	// there is no inheritance in GO lang
	// there is no concepts like super or Parent
	//child sild ni kei hunna mya
	syrus := User{"Syrus", "Syrus@go.dev", true, 16}
	fmt.Println(syrus)
	fmt.Printf("Syrus Details are: %+v", syrus)                   // structure lai %+v ho hai bro
	fmt.Printf("Name is %v Email is %v", syrus.Name, syrus.Email) // structure lai %+v ho hai bro

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

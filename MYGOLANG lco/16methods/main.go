package main

import "fmt"

func main() {

	fmt.Println("Structs in goLang")

	// there is no inheritance in GO lang
	// there is no concepts like super or Parent
	//child sild ni kei hunna mya
	syrus := User{"Syrus", "Syrus@go.dev", true, 16}
	fmt.Println(syrus)
	fmt.Printf("Syrus Details are: %+v", syrus)                     // structure lai %+v ho hai bro
	fmt.Printf("Name is %v Email is %v\n", syrus.Name, syrus.Email) // structure lai %+v ho hai bro
	syrus.GetSTatus()
	syrus.NewMail()
	fmt.Printf("Name is %v Email is %v\n", syrus.Name, syrus.Email) // change gardaina email kina bhane copoy banaucha tei bhayera pointer user garnu parcha
	syrus.NewAge()
	syrus.NewUser()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetSTatus() {
	fmt.Println("Is user active:", u.Status)

}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}

func (u User) NewAge() {
	u.Age = 18
	fmt.Printf("You are %v years old\n", u.Age)
}
func (u User) NewUser() {
	u.Name = "Sy"
	fmt.Println(u.Name)

}

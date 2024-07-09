package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File")
	content := "This needs to go on file"
	file, err := os.Create("./mysyrusgofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is", length)
	defer file.Close()
	readFile("./mysyrusgofile.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text dats inside file is \n", string(databyte))

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

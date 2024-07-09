package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T", response)

	defer response.Body.Close() // caller's responsibility to close connection
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// content := string(databytes)
	// fmt.Println(content)
	fmt.Println(string(databytes))
}

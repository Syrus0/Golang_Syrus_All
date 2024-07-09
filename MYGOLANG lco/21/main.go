package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verb video")
	//PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content) //content byte format ma cha la bro
	// fmt.Println(string(content))
}
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
{
"coursename" : "Let's go with Golang",
"price" : 0,
"platform" : "learncodeonline.in"

}


`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("firstname", "syrus")
	data.Add("lastname", "rajbhandari")
	data.Add("email", "syrus.go.dev")
	data.Add("age", "22")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome, Go Server")
	GetHelper()
}

// Function to Generate a URL
func GenerateUrl(requestType string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   "localhost:8000",
		Path:   requestType}
}

func GetHelper() {
	// GET URL
	myGetUrl := GenerateUrl("get")

	fmt.Println("My Url -", myGetUrl)
	fmt.Printf("Type of URL - %T\n", myGetUrl)

	// Make Http, Get Request
	response, err := http.Get(myGetUrl.String())

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code -", response.StatusCode)
	fmt.Println("Content Length -", response.ContentLength)

	// Read Content of Response
	// One way to use strings.Builder
	var responseString strings.Builder

	// Read the Content of response, use ioutil
	// Inevitable line
	content, _ := ioutil.ReadAll(response.Body)

	// 1st way continued
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count -", byteCount)
	fmt.Println("Data -", responseString.String())

	// Other way to read content
	// fmt.Println("Data -", string(content))

}

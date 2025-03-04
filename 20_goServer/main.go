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
	PostJsonHelper()
	PostFormHandler()
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

func PostJsonHelper() {

	// POST Url
	postUrl := GenerateUrl("post")

	fmt.Println("My Url -", postUrl)
	fmt.Printf("Type of URL - %T\n", postUrl)

	// Fake JSON Data
	requestBody := strings.NewReader(`
		{
			"coursename": "Golang",
			"price": 0,
			"org": "XoXo"
		}
	`)

	response, err := http.Post(postUrl.String(), "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content = ", string(content))
}

func PostFormHandler() {
	// URL
	postFormUrl := GenerateUrl("postform")

	fmt.Println("My Url -", postFormUrl)
	fmt.Printf("Type of URL - %T\n", postFormUrl)

	// Fake Form Data
	data := url.Values{}
	data.Add("firstName", "Dhairya")
	data.Add("lastName", "Hans")
	data.Add("clan", "XoXo")

	response, err := http.PostForm(postFormUrl.String(), data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content - ", string(content))

}

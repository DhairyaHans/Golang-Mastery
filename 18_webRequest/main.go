package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response - %T\n", response)
	fmt.Println("Response -", response)

	defer response.Body.Close() // Caller's responsibilty to Close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	data := string(dataBytes)
	fmt.Print("Response Body - %v", data)

}

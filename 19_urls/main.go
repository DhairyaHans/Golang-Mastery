package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome, Handling URLs")

	// Parse the URL
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of result - %T\n", result)

	// Parts of URL
	// fmt.Println("URL Scheme", result.Scheme)
	// fmt.Println("URL Host", result.Host)
	// fmt.Println("URL Path", result.Path)
	// fmt.Println("URL Port", result.Port())
	fmt.Println("URL Raw Query", result.RawQuery)

	qParams := result.Query()

	for key, val := range qParams {
		fmt.Printf("Value of param, %v is %v\n", key, val)
	}

	// Creating URL
	// Must add &
	reconstruct_url := &url.URL{
		Scheme:   result.Scheme,
		Host:     result.Host,
		Path:     result.Path,
		RawQuery: result.RawQuery}

	anotherUrl := reconstruct_url.String()
	fmt.Println("New Url -", anotherUrl)
}

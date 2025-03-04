package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// Create Alias for the Variables, using ``
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // It will remove the Password field from json
	Tags     []string `json:"tags,omitempty"` // Dont show tags, if empty
}

func main() {
	fmt.Println("Welcome, JSON Details")
	// EncodeJson()
	DecodeJson()
}

// Create JSON from a struct
func EncodeJson() {
	courseList := []course{
		{"ReactJs", 299, "gulugulu.com", "abc123", []string{"Web-dev", "Js"}},
		{"MERN", 799, "gulugulu.com", "xlpop123", []string{"Full-Stack", "Js"}},
		{"Angular", 299, "gulugulu.com", "xxx123", nil},
	}

	fmt.Println(courseList)
	// Package this data as JSON data
	jsonData, err := json.Marshal(courseList)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	// Package this data as JSON data in an Indented form
	jsonDataIndented, err := json.MarshalIndent(courseList, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonDataIndented)

}

// Consume/Read JSON Data
func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "MERN",
		"Price": 799,
		"website": "gulugulu.com",
		"tags": ["Full-Stack","Js"]
	}
	`)

	// Storing JSON data in a struct
	var myCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON is not Valid")
	}

	// Storing the JSON data in the form of key-value pairs
	var myJsonDt map[string]interface{}
	json.Unmarshal(jsonData, &myJsonDt)
	fmt.Printf("%#v\n", myJsonDt)

	for k, v := range myJsonDt {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}
}

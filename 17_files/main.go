package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filename string) {
	// Whenever you read or fetch data from a file or web... its always in Bytes format
	dataByte, err := ioutil.ReadFile(filename)
	CheckNilErr(err)

	fmt.Printf("File content in bytes -%v\n", dataByte)
	fmt.Printf("Actual file content -%v\n", string(dataByte))

}

func main() {
	fmt.Println("Welcome, Files")

	content := "Hello there, this is your first file in Go"

	file, err := os.Create("./firstFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	CheckNilErr(err)

	length, err := io.WriteString(file, content)
	CheckNilErr(err)

	fmt.Println("Length of file -", length)

	defer file.Close()

	ReadFile("./firstFile.txt")

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhairyahans/mongoapi/routers"
)

func main() {
	fmt.Println("Welcome to Mongo APIs")
	fmt.Println("Server is getting started...")
	router := routers.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listening on Port 4000")
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to, GO Mod(Modules)")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hello from Go MOD")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<H1>Welcome to GO MODs... Hola Amigos Gracias</H1>"))
}

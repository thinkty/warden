package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")
	http.Handle("/", http.FileServer(http.Dir("./"))) // TODO: only allow index.html

	log.Fatal(http.ListenAndServe(":8080", nil))
}

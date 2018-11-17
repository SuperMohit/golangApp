package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Golang App version 1")
		log.Println("Hello Mars!")
	}))
	log.Println("Now server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}

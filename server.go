package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HandleFunc function to add route handlers to the web serve
	// ResponseWriter to send a response back
	// Request object that provides more information about the request itself
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	fmt.Println("server starting at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

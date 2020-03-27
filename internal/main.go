package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(responseWriter, "Hello, %q", html.EscapeString(request.URL.Path))
	})
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

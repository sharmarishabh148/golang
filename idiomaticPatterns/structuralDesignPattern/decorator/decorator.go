package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Add behavior,
without modifying original function/method,
keeping same interface.
Useful for adhering to the Single Responsibility Principle
*/

func main() {
	http.HandleFunc("/", helloEndpoint)
	http.HandleFunc("/logged", logDecorator(helloEndpoint))
	_ = http.ListenAndServe(":8080", nil)
}

func helloEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func logDecorator(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		fn(w, r)
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now().Format("2006/01/02 15:04:05"))
	fmt.Println("hello, world")
}


func greets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now().Format("2006/01/02 15:04:05"))
	fmt.Println("hello, world-get")
}
func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/get", greets)
	http.ListenAndServe(":8333", nil)
}
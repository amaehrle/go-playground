package main

import (
	"io"
	"net/http"
  "fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
  fmt.Println("starting service on port 8888")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}

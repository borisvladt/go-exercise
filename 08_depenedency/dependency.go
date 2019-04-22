package main

import (
	"fmt"
	"io"
	"net/http"
	//"os"
)

func Greet(writer io.Writer, name string, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %s\n %v", name, request)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world", r)
}

func main() {
	//Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

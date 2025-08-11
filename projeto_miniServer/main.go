package main

import (
	"fmt"
	"net/http"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	http.HandleFunc("/", handlerHello)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)

}

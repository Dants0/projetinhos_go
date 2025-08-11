package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerHello)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}
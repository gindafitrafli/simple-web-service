package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "time"
)
func main() {
	fmt.Println("hello world")
	r := mux.NewRouter()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
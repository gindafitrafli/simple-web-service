package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	controller "github.com/gindafitrafli/go/simple-web-service/api/muxnethttp"
	// "time"
)
func main() {
	fmt.Println("hello world")
	//using mux and pure net/http

	r := mux.NewRouter()
	r.HandleFunc("/outlet", controller.CreateOutlet).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))

	//using gin
	
}
package main

import (
	"fmt"
	"net/http"

	"github.com/schlucht/rechnung/pkg/handlers"
)

const PORTNUMBER = ":8080"



func main() {
	

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	
	fmt.Println("Server run on localhost:8080")
	_ = http.ListenAndServe(PORTNUMBER, nil)

}

package main

import (
	"log"
)



func main() {
	
	usr := User {
		FirstName: "Lothar",
		LastName: "Schmid",
		Age: 55,
	}
	
}

func saySomething(s3 string) (string, string) {
	log.Println("this s3 come from func", s3)
	return s3, "World"
}

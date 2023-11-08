package main

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))

}
package main

import (
	"api/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting API...")
	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":5000", r))
}

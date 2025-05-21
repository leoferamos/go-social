package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/routes"
)

func main() {
	fmt.Print("Running web application...")
	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":8080", r))
}

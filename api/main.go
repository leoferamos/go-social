package main

import (
	"api/src/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":5000", r))
}

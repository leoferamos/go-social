package main

import (
	"fmt"
	"go_social/internal/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting API...")
	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":5000", r))
}

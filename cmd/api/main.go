package main

import (
	"fmt"
	"go_social/config"
	"go_social/internal/routes"
	"log"
	"net/http"
)

func main() {

	config.Init()
	fmt.Println("Starting server on port:", config.Port)

	r := routes.Generate()
	fmt.Println("Router generated successfully")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

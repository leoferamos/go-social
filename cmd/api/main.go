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
	fmt.Println(config.Port)
	r := routes.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

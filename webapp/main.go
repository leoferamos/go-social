package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/routes"
	"webapp/src/utils"
)

func main() {
	fmt.Print("Running web application...")
	utils.LoadTemplates()

	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":8080", r))
}

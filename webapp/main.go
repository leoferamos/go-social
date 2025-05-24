package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webapp/src/config"
	"webapp/src/routes"
	"webapp/src/utils"
)

func main() {
	config.Load()
	fmt.Print("Running web application...")
	utils.LoadTemplates()

	r := routes.Generate()

	port := os.Getenv("WEBAPP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}

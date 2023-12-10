package main

import (
	"go-server/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.Router()
	log.Println("Starting server.")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server Runs at port 8000")
}

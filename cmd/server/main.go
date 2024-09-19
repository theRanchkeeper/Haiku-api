package main

import (
	"api/cmd/internals/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.CreateMux()

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err.Error())
	}
}

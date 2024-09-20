package routes

import (
	"fmt"
	"net/http"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Add all the routes here
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hello world!! Server is up\n")
	})

	return mux
}

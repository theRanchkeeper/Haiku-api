package routes

import (
	"fmt"
	"net/http"
)

func getSimpleData(w http.ResponseWriter, r *http.Request) {
	fmt.Print("data from simple")
}

package handlers

import (
	"fmt"
	"log"
	"net/http"

	"harvest/controllers"
)

func RegisterIndex() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Received request on /")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	name := r.URL.Query().Get("name")
	fmt.Fprint(w, controllers.Greeting(name))
}

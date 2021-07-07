package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"harvest/controllers"
)

type indexResult struct {
	result string
}

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
	greeting := controllers.Greeting(name)
	b, err := json.Marshal(indexResult{result: greeting})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(b)
}

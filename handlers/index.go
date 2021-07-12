package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"harvest/controllers"
	"harvest/core"
)

type indexResult struct {
	result string
}

func RegisterIndex() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	trace := core.GetTraceId(r.Header.Get("X-Cloud-Trace-Context"))
	ctx := context.WithValue(context.Background(), "trace", trace)

	core.Info(ctx, "Default log", "indexHandler")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	name := r.URL.Query().Get("name")
	greeting := controllers.Greeting(name)
	b, err := json.Marshal(indexResult{result: greeting})
	if err != nil {
		core.Error(ctx, "Failed to marshal response object", "indexHandler")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(b)
}

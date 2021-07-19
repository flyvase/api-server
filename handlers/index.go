package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"harvest/logger"
	"harvest/repositories"
)

type IndexResponseBody struct {
	Message string `json:"message"`
}

func RegisterIndex() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	const component = "indexHandler"
	trace := GetTraceId(r.Header.Get("X-Cloud-Trace-Context"))
	ctx := context.WithValue(context.Background(), "trace", trace)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	db, err1 := repositories.InitMySqlConnection()
	if err1 != nil {
		logger.Error(ctx, err1.Error(), component)
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	if err := db.Ping(); err != nil {
		logger.Error(ctx, err.Error(), component)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(IndexResponseBody{Message: "Database connection confirmed"})
	if err != nil {
		logger.Error(ctx, err.Error(), component)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(b)
}

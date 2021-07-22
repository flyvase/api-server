package handlers

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"harvest/config"
	"harvest/logger"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	const component = "usersHandler"
	logger.Debug(fmt.Sprintf("url: %v, method: %v", r.URL, r.Method), component)
}

func UsersHandler() http.Handler {
	handler := http.HandlerFunc(usersHandler)

	return cors.New(cors.Options{
		AllowedOrigins: []string{config.AllowedOrigin()},
		AllowedMethods: []string{"POST"},
	}).Handler(handler)
}

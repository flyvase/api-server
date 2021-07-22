package handlers

import (
	"context"
	"net/http"

	"github.com/rs/cors"

	"harvest/config"
	"harvest/logger"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	const component = "usersHandler"
	trace := getTraceId(r)
	ctx := context.WithValue(context.Background(), "trace", trace)

	validatePath("/users/", w, r)
	validateMethods([]string{http.MethodPost}, w, r)

	logger.Info(ctx, "good", component)
}

func UsersHandler() http.Handler {
	handler := http.HandlerFunc(usersHandler)

	return cors.New(cors.Options{
		AllowedOrigins: []string{config.AllowedOrigin()},
		AllowedMethods: []string{http.MethodPost},
	}).Handler(handler)
}

package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/cors"

	"harvest/config"
	"harvest/entities"
	"harvest/logger"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	const component = "usersHandler"
	trace := getTraceId(r)
	ctx := context.WithValue(context.Background(), "trace", trace)

	validatePath("/users/", w, r)
	validateMethods([]string{http.MethodPost}, w, r)
	validateContentType(jsonContentType, w, r)

	var b entities.User
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&b); err != nil {
		logger.Error(ctx, component, errors.WithStack(err))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

}

func UsersHandler() http.Handler {
	handler := http.HandlerFunc(usersHandler)

	return cors.New(cors.Options{
		AllowedOrigins: []string{config.AllowedOrigin()},
		AllowedMethods: []string{http.MethodPost},
	}).Handler(handler)
}

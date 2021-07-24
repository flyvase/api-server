package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"github.com/rs/cors"

	"harvest/config"
	"harvest/controllers"
	"harvest/core"
	"harvest/entities"
	"harvest/interfaces"
	"harvest/logger"
)

func usersHandler(w http.ResponseWriter, r *http.Request, i interfaces.User) {
	const component = "usersHandler"
	trace := getTraceId(r)
	ctx := context.WithValue(context.Background(), "trace", trace)

	validatePath("/users/", w, r)
	validateMethods([]string{http.MethodPost}, w, r)
	validateContentType(jsonContentType, w, r)

	var u entities.User
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&u); err != nil {
		logger.Error(ctx, component, errors.WithStack(err))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	v := validator.New()
	if err := v.Struct(u); err != nil {
		logger.Error(ctx, component, errors.WithStack(err))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := controllers.CreateUser(i, u); err != nil {
		logger.Error(ctx, component, err)
		switch err.(type) {
		case core.DSConnErr:
			http.Error(w, "Data source unavailable", http.StatusInternalServerError)
		case core.UnknownErr:
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		default:
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UsersHandler(i interfaces.User) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usersHandler(w, r, i)
	})

	return cors.New(cors.Options{
		AllowedOrigins: []string{config.AllowedOrigin()},
		AllowedMethods: []string{http.MethodPost},
	}).Handler(handler)
}

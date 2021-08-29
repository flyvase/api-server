package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"

	"harvest/controllers"
	"harvest/core"
	"harvest/entities"
	"harvest/interfaces"
	"harvest/logger"
)

func usersHandler(w http.ResponseWriter, r *http.Request, provider *interfaces.RepositoriesProvider) {
	const component = "UsersHandler"
	trace := getTraceId(r)

	var u entities.User
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&u); err != nil {
		logger.Error(component, err, trace)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	v := validator.New()
	if err := v.Struct(u); err != nil {
		logger.Error(component, err, trace)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := controllers.CreateUser(provider.User, u); err != nil {
		logger.Error(component, err, trace)
		switch err.(type) {
		case core.DSConnErr:
			http.Error(w, "Data source unavailable", http.StatusInternalServerError)
		default:
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UsersHandler(provider *interfaces.RepositoriesProvider) http.Handler {
	opt := HandlerOptions{
		Path:        "users/",
		Methods:     &[]string{http.MethodPost},
		ContentType: jsonContentType,
	}
	return NewHandler(&opt, usersHandler, provider)
}

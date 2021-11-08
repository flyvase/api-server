package handler

import (
	"net/http"

	"harvest/src/controller"
	"harvest/src/core/apperror"
	"harvest/src/core/logger"
	"harvest/src/domain/repository"
	"harvest/src/infrastructure/http/middleware"
	"harvest/src/infrastructure/http/request"
)

func UsersPost(authR repository.Auth, userR repository.User) http.Handler {
	return middleware.DefaultPostMiddlewares(
		authR,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const component = "UsersPostHandler"
			trace := request.GetTraceId(r)

			var body request.User
			if err := request.DecodeUsersPostRequest(r.Body, &body); err != nil {
				logger.Error(component, err, trace)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			entity := body.ToUserEntity()

			if err := controller.CreateUser(entity, userR, authR); err != nil {
				logger.Error(component, err, trace)
				switch err.(type) {
				case apperror.SqlConnClosed:
					http.Error(w, "Database is not available", http.StatusInternalServerError)
				default:
					http.Error(w, "Unknown error", http.StatusInternalServerError)
				}
				return
			}

			w.WriteHeader(http.StatusCreated)
		}),
	)
}

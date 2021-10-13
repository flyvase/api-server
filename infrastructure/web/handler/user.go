package handler

import (
	"net/http"

	"harvest/controller"
	"harvest/core/apperror"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/web/middleware"
	"harvest/infrastructure/web/request"
)

func UserPost(authR repository.Auth, userR repository.User) http.Handler {
	return middleware.DefaultPostMiddlewares(
		authR,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if path != "/user/" {
				http.NotFound(w, r)
				return
			}

			const component = "UserPostHandler"
			trace := request.GetTraceId(r)

			var body request.User
			if err := request.DecodeUserRequestJson(r.Body, &body); err != nil {
				logger.Error(component, err, trace)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			entity := body.ToUserEntity()

			if err := controller.CreateUser(entity, userR, authR); err != nil {
				logger.Error(component, err, trace)
				switch err.(type) {
				case apperror.SqlConnClosed:
					http.Error(w, "Datasource unavailable", http.StatusInternalServerError)
				default:
					http.Error(w, "Unknown error", http.StatusInternalServerError)
				}
				return
			}

			w.WriteHeader(http.StatusCreated)
		}),
	)
}

package handler

import (
	"net/http"

	"harvest/controller"
	"harvest/core/apperror"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/web/request"
)

const uhComponent = "UserHandler"

func userHandler(userR repository.User, authR repository.Auth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trace := request.GetTraceId(r)

		var ur request.User
		if err := request.DecodeUserRequestJson(r.Body, &ur); err != nil {
			logger.Error(uhComponent, err, trace)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ue := ur.ToUserEntity()

		if err := controller.CreateUser(ue, userR, authR); err != nil {
			logger.Error(uhComponent, err, trace)
			switch err.(type) {
			case apperror.SqlConnClosed:
				http.Error(w, "Datasource unavailable", http.StatusInternalServerError)
			case apperror.Unknown:
				http.Error(w, "Unknown error", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

func UserHandler(userR repository.User, authR repository.Auth) http.Handler {
	opt := Option{
		Path:        "/user/",
		Methods:     &[]string{http.MethodPost},
		ContentType: jsonContentType,
	}
	return buildHandlerWithDefaultMiddlewares(&opt, userHandler(userR, authR), authR)
}

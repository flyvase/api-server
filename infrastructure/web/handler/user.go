package handler

import (
	"net/http"

	"harvest/controller"
	"harvest/core/exception"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/web/request"
)

const uhComponent = "UserHandler"

func userHandler(usr repository.User) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trace := request.GetTraceId(r)

		var ur request.User
		if err := request.DecodeUserRequestJson(r.Body, &ur); err != nil {
			logger.Error(uhComponent, err, trace)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ue := ur.ToUserEntity()

		if err := controller.CreateUser(ue, usr); err != nil {
			logger.Error(uhComponent, err, trace)
			switch err.(type) {
			case exception.SqlConnClosedError:
				http.Error(w, "Datasource unavailable", http.StatusInternalServerError)
			case exception.UnknownError:
				http.Error(w, "Unknown error", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

func UserHandler(usr repository.User, aur repository.Auth) http.Handler {
	opt := Option{
		Path:        "/user/",
		Methods:     &[]string{http.MethodPost},
		ContentType: jsonContentType,
	}
	return buildHandlerWithDefaultMiddlewares(&opt, userHandler(usr), aur)
}

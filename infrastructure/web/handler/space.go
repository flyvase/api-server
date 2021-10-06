package handler

import (
	"net/http"

	"harvest/controller"
	"harvest/core/apperror"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/web/request"
	"harvest/infrastructure/web/response"
)

const shComponent = "SpaceHandler"

func spaceHandler(spaceR repository.Space) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trace := request.GetTraceId(r)

		if spaces, err := controller.FetchSpace(spaceR); err != nil {
			logger.Error(shComponent, err, trace)
			switch err.(type) {
			case apperror.SqlConnClosed:
				http.Error(w, "Datasource unavailable", http.StatusInternalServerError)
			case apperror.Unknown:
				http.Error(w, "Unknown error", http.StatusInternalServerError)
			}
			return
		} else {
			response.MarshalSpaceResponseJson(w, spaces)
		}

		w.WriteHeader(http.StatusCreated)
	})
}

func SpaceHandler(SpaceR repository.Space, authR repository.Auth) http.Handler {
	opt := Option{
		Path:        "/space/",
		Methods:     &[]string{http.MethodGet},
		ContentType: jsonContentType,
	}
	return buildHandlerWithDefaultMiddlewares(&opt, spaceHandler(SpaceR), authR)
}

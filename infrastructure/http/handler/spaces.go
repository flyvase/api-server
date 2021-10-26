package handler

import (
	"net/http"

	"harvest/controller"
	"harvest/core/apperror"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/http/middleware"
	"harvest/infrastructure/http/request"
	"harvest/infrastructure/http/response"
)

func SpacesGet(authR repository.Auth, spaceR repository.Space) http.Handler {
	return middleware.DefaultGetMiddlewares(
		authR,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const component = "SpacesGetHandler"
			trace := request.GetTraceId(r)

			spaces, err := controller.ListSpaces(spaceR)
			if err != nil {
				logger.Error(component, err, trace)
				switch err.(type) {
				case apperror.SqlConnClosed:
					http.Error(w, "Database is not available", http.StatusInternalServerError)
				default:
					http.Error(w, "Unknown error", http.StatusInternalServerError)
				}
				return
			}

			json, err := response.EncodeSpaceEntities(spaces)

			if err != nil {
				logger.Error(component, err, trace)
				http.Error(w, "Failed to encode space data to json", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(json)
		}),
	)
}

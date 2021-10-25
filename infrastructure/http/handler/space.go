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

func SpaceGet(authR repository.Auth, spaceR repository.Space) http.Handler {
	return middleware.DefaultGetMiddlewares(
		authR,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if path != "/space/" {
				http.NotFound(w, r)
				return
			}

			const component = "SpaceGetHandler"
			trace := request.GetTraceId(r)

			spaces, err := controller.FetchSpace(spaceR)
			if err != nil {
				logger.Error(component, err, trace)
				switch err.(type) {
				case apperror.SqlConnClosed:
					http.Error(w, "Datasource unavailable", http.StatusInternalServerError)
				default:
					http.Error(w, "Unknown error", http.StatusInternalServerError)
				}
				return
			}

			json, err := response.EncodeSpaceEntities(spaces)

			if err != nil {
				logger.Error(component, err, trace)
				http.Error(w, "Json encode error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(json)
		}),
	)
}

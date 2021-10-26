package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func SpaceDetailsGet(authR repository.Auth, spaceR repository.Space) http.Handler {
	return middleware.DefaultGetMiddlewares(
		authR,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const component = "SpaceDetailsGetHandler"
			trace := request.GetTraceId(r)

			vars := mux.Vars(r)
			id, err := strconv.ParseUint(vars["space_id"], 10, 32)
			if err != nil {
				logger.Error(component, err, trace)
				http.Error(w, "Invalid space id", http.StatusBadRequest)
				return
			}

			space, err := controller.FetchSpace(uint32(id), spaceR)
			if err != nil {
				logger.Error(component, err, trace)
				switch err.(type) {
				case apperror.SqlConnClosed:
					http.Error(w, "Database is not available", http.StatusInternalServerError)
				case apperror.EmptySqlResult:
					http.Error(w, "No matching space with provided id", http.StatusBadRequest)
				default:
					http.Error(w, "Unknown error", http.StatusInternalServerError)
				}
				return
			}

			json, err := response.EncodeSpaceEntity(space)
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

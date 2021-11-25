package handler

import (
	"harvest/src/core/errors"
	"harvest/src/core/logger"
	"harvest/src/domain/repository"
	"harvest/src/domain/value"
	"harvest/src/infrastructure/http/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SpaceDetailsGet(spaceRepository repository.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		const component = "spaceDetailsGetHandler"
		trace := getTraceId(r)

		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["space_id"], 10, 32)
		if err != nil {
			http.Error(rw, "Invalid space id", http.StatusBadRequest)
			return
		}

		spaceModel, err := spaceRepository.Fetch(
			value.SpaceId{
				Value: uint(id),
			},
		)
		if err != nil {
			switch err.(type) {
			case *errors.Unexpected:
				{
					logger.Error(component, err, trace)
					http.Error(rw, "Unexpected error", http.StatusInternalServerError)
					return
				}
			default:
				{
					logger.Error(component, err, trace)
					http.Error(rw, "Data source unavailable", http.StatusInternalServerError)
					return
				}
			}
		}

		if spaceModel == nil {
			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		json, err := response.EncodeSpaceModel(spaceModel)
		if err != nil {
			logger.Error(component, err, trace)
			http.Error(rw, "Failed to encode resources", http.StatusInternalServerError)
			return
		}

		respond(rw, http.StatusOK, json)
	})
}
package handler

import (
	"api-server/src/core/errors"
	"api-server/src/core/logger"
	"api-server/src/domain/repository"
	"api-server/src/infrastructure/http/response"
	"net/http"
)

func SpacesGet(spaceRepository repository.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		const component = "spacesGetHandler"
		trace := getTraceId(r)

		spaceModels, err := spaceRepository.List()
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

		json, err := response.EncodeSpaceModels(spaceModels)
		if err != nil {
			logger.Error(component, err, trace)
			http.Error(rw, "Failed to encode resources", http.StatusInternalServerError)
			return
		}

		respond(rw, http.StatusOK, json)
	})
}

package handler

import (
	"harvest/src/adapter/http/response"
	"harvest/src/domain/repository"
	"net/http"
)

func SpacesGet(spaceRepository repository.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		spaceModels, _ := spaceRepository.List()

		json, _ := response.EncodeSpaceModels(spaceModels)

		respond(rw, http.StatusOK, json)
	})
}

package handler

import (
	"harvest/src/adapter/http/response"
	"harvest/src/application/usecase"
	"net/http"
)

func SpacesGet(spaceInteractor usecase.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		spaceModels, _ := spaceInteractor.List()

		json, _ := response.EncodeSpaceModels(spaceModels)

		respond(rw, http.StatusOK, json)
	})
}

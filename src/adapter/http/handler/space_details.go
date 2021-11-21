package handler

import (
	"harvest/src/adapter/http/response"
	"harvest/src/application/usecase"
	"harvest/src/domain/value"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SpaceDetailsGet(spaceInteractor usecase.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["space_id"], 10, 32)
		if err != nil {
			http.Error(rw, "Invalid space id", http.StatusBadRequest)
			return
		}

		spaceModel, _ := spaceInteractor.Fetch(
			value.SpaceId{
				Value: uint(id),
			},
		)

		json, _ := response.EncodeSpaceModel(spaceModel)

		respond(rw, http.StatusOK, json)
	})
}

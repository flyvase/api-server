package handler

import (
	"harvest/src/adapter/http/response"
	"harvest/src/domain/repository"
	"harvest/src/domain/value"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SpaceDetailsGet(spaceRepository repository.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["space_id"], 10, 32)
		if err != nil {
			http.Error(rw, "Invalid space id", http.StatusBadRequest)
			return
		}

		spaceModel, _ := spaceRepository.Fetch(
			value.SpaceId{
				Value: uint(id),
			},
		)

		json, _ := response.EncodeSpaceModel(spaceModel)

		respond(rw, http.StatusOK, json)
	})
}

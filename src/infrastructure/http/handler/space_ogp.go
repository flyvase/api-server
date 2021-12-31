package handler

import (
	"encoding/json"
	"harvest/src/core/errors"
	"harvest/src/core/logger"
	"harvest/src/domain/repository"
	"harvest/src/domain/value"
	"net/http"
	"strconv"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/gorilla/mux"
)

func SpaceOgpGet(spaceRepository repository.Space) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		const component = "spaceOgpGetHandler"
		trace := getTraceId(r)

		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["space_id"], 10, 32)
		if err != nil {
			http.Error(rw, "Invalid space id", http.StatusBadRequest)
			return
		}

		url, err := spaceRepository.GetWebsiteUrl(
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

		if url == "" {
			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			logger.Error(component, err, trace)
			http.Error(rw, http.StatusText(resp.StatusCode), resp.StatusCode)
			return
		}

		ogp := opengraph.NewOpenGraph()
		if err := ogp.ProcessHTML(resp.Body); err != nil {
			logger.Error(component, err, trace)
			http.Error(rw, "Failed to parse HTML", http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(ogp)
		if err != nil {
			logger.Error(component, err, trace)
			http.Error(rw, "Failed to encode resources", http.StatusInternalServerError)
			return
		}

		respond(rw, http.StatusOK, js)
	})
}
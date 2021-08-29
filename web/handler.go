package web

import (
	"net/http"

	"github.com/rs/cors"

	"harvest/config"
	"harvest/interfaces"
)

const jsonContentType = "application/json"

type HandlerOptions struct {
	Path        string
	Methods     *[]string
	ContentType string
}

type Handler func(w http.ResponseWriter, r *http.Request, provider *interfaces.RepositoriesProvider)

func NewHandler(opt *HandlerOptions, handler Handler, provider *interfaces.RepositoriesProvider) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigin(),
		AllowedMethods: *opt.Methods,
	}).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validateRequest(opt, w, r)
		handler(w, r, provider)
	}))
}

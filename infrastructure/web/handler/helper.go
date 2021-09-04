package handler

import (
	"net/http"

	"harvest/domain/repository"
	"harvest/infrastructure/web/middleware"
)

const jsonContentType = "application/json"

type Option struct {
	Path        string
	Methods     *[]string
	ContentType string
}

func buildHandlerWithDefaultMiddlewares(opt *Option, next http.Handler, aur repository.Auth) http.Handler {
	return middleware.Logger(
		middleware.Auth(
			middleware.PathValidator(
				middleware.MethodValidator(
					middleware.ContentTypeValidator(
						next, opt.ContentType,
					),
					*opt.Methods,
				),
				opt.Path,
			),
			aur,
		),
	)
}

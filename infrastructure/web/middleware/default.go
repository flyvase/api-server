package middleware

import (
	"net/http"

	"harvest/domain/repository"
)

func DefaultGetMiddlewares(authR repository.Auth, next http.Handler) http.Handler {
	return requestLogger(
		auth(
			authR,
			next,
		),
	)
}

func DefaultPostMiddlewares(authR repository.Auth, next http.Handler) http.Handler {
	return requestLogger(
		auth(
			authR,
			contentTypeValidator(
				"application/json",
				next,
			),
		),
	)
}

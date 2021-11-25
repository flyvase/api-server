package middleware

import (
	"harvest/src/domain/repository"
	"net/http"
)

func Defaults(authRepository repository.Auth, handler http.Handler) http.Handler {
	return Logger(
		Auth(
			authRepository,
			handler,
		),
	)
}

package middleware

import (
	"net/http"
	"strings"

	"harvest/core/exception"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/web/request"
)

const aumComponent = "AuthMiddleware"

func Auth(next http.Handler, auth repository.Auth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token := strings.Fields(authHeader)[1]
		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if err := auth.VerifyToken(token); err != nil {
			trace := request.GetTraceId(r)

			switch err.(type) {
			case exception.UnknownError:
				logger.Error(aumComponent, err, trace)
			}

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
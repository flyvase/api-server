package middleware

import (
	"net/http"
	"strings"

	"harvest/core/apperror"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/http/request"
)

func auth(authR repository.Auth, next http.Handler) http.Handler {
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

		if err := authR.VerifyToken(token); err != nil {
			trace := request.GetTraceId(r)

			switch err.(type) {
			case apperror.Unknown:
				logger.Error("AuthMiddleware", err, trace)
			}

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

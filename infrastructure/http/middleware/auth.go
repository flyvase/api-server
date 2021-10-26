package middleware

import (
	"net/http"
	"strings"

	"harvest/controller"
	"harvest/core/apperror"
	"harvest/core/logger"
	"harvest/domain/repository"
	"harvest/infrastructure/http/request"
)

func auth(authR repository.Auth, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		token := strings.Fields(authHeader)[1]
		if token == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		if err := controller.VerifyAuthToken(token, authR); err != nil {
			trace := request.GetTraceId(r)

			switch err.(type) {
			case apperror.Unknown:
				logger.Error("AuthMiddleware", err, trace)
				http.Error(w, "Unknown error", http.StatusUnauthorized)
				return
			}

			// Want to provide detailed message but the handling will be complicated
			// Tmp solution
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

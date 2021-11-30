package middleware

import (
	"harvest/src/domain/repository"
	"net/http"
)

func Auth(authRepository repository.Auth, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// authHeader := r.Header.Get("Authorization")
		// if authHeader == "" {
		// 	http.Error(rw, "Authorization header is required", http.StatusUnauthorized)
		// 	return
		// }

		// split := strings.Fields(authHeader)
		// if len(split) != 2 {
		// 	http.Error(rw, "Invalid Authorization header format", http.StatusUnauthorized)
		// 	return
		// }

		// token := split[1]
		// if token == "" {
		// 	http.Error(rw, "Authorization token is required", http.StatusUnauthorized)
		// 	return
		// }

		// if err := authRepository.VerifyToken(token); err != nil {
		// 	http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		// 	return
		// }

		next.ServeHTTP(rw, r)
	})
}

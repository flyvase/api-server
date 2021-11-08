package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"harvest/src/config"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.Mode == "debug" {
			// https://pkg.go.dev/time#pkg-constants
			msg := fmt.Sprintf("[%v] path: %s, method: %s", time.Now().Format("15:04:05"), r.URL.Path, r.Method)
			log.Println(msg)
		}

		next.ServeHTTP(w, r)
	})
}

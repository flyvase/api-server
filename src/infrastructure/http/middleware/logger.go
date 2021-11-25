package middleware

import (
	"fmt"
	"harvest/src/config"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if config.Mode == "debug" {
			now := time.Now()
			msg := fmt.Sprintf(
				"%s %s %s %s %s",
				now.Format("2006-01-02"),
				now.Format("15:04:05"),
				r.Method,
				r.Proto,
				r.URL.Path,
			)
			log.Println(msg)
		}

		next.ServeHTTP(rw, r)
	})
}

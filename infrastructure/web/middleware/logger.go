package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"harvest/config"
)

const timeStamp = "15:04:05"

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.Mode == "debug" {
			msg := fmt.Sprintf("[%v] path: %s, method: %s", time.Now().Format(timeStamp), r.URL.Path, r.Method)
			log.Println(msg)
		}

		next.ServeHTTP(w, r)
	})
}

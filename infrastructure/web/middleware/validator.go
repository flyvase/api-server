package middleware

import "net/http"

func PathValidator(next http.Handler, path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MethodValidator(next http.Handler, methods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flag := false
		for _, m := range methods {
			if r.Method == m {
				flag = true
			}
		}

		if !flag {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ContentTypeValidator(next http.Handler, contentType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != contentType {
			http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}

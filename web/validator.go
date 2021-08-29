package web

import (
	"fmt"
	"net/http"
)

func isValidPath(path string, actualPath string) bool {
	return actualPath == path
}

func isValidMethods(methods []string, actualMethod string) bool {
	flag := false
	for _, m := range methods {
		if actualMethod == m {
			flag = true
		}
	}

	return flag
}

func isValidContentType(ct string, actualCt string) bool {
	return actualCt == ct
}

func validateRequest(opt *HandlerOptions, w http.ResponseWriter, r *http.Request) {
	if isValidPath(opt.Path, r.URL.Path) {
		http.NotFound(w, r)
	}

	if isValidMethods(*opt.Methods, r.Method) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	ct := r.Header.Get("Content-Type")
	if isValidContentType(opt.ContentType, ct) {
		msg := fmt.Sprintf("Invalid content type: %s", ct)
		http.Error(w, msg, http.StatusUnsupportedMediaType)
	}
}

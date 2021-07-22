package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"harvest/config"
)

func getTraceId(r *http.Request) string {
	parts := strings.Split(r.Header.Get("X-Cloud-Trace-Context"), "/")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", config.ProjectId, parts[0])
	}
	return ""
}

func validatePath(p string, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != p {
		http.NotFound(w, r)
	}
}

func validateMethods(methods []string, w http.ResponseWriter, r *http.Request) {
	for _, m := range methods {
		if r.Method == m {
			return
		}
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

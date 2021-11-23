package handler

import (
	"fmt"
	"harvest/src/config"
	"net/http"
	"strings"
)

func respond(rw http.ResponseWriter, code int, body []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(body)
}

func getTraceId(r *http.Request) string {
	parts := strings.Split(r.Header.Get("X-Cloud-Trace-Context"), "/")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", config.ProjectId, parts[0])
	}
	return ""
}

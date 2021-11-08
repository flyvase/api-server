package request

import (
	"fmt"
	"net/http"
	"strings"

	"harvest/src/config"
)

func GetTraceId(r *http.Request) string {
	parts := strings.Split(r.Header.Get("X-Cloud-Trace-Context"), "/")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", config.ProjectId, parts[0])
	}
	return ""
}

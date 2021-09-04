package handler

import (
	"fmt"
	"net/http"
	"strings"

	"harvest/config"
	"harvest/infrastructure/web/middleware"
)

const jsonContentType = "application/json"

type Option struct {
	Path        string
	Methods     *[]string
	ContentType string
}

func getTraceId(r *http.Request) string {
	parts := strings.Split(r.Header.Get("X-Cloud-Trace-Context"), "/")
	if len(parts) > 0 && len(parts[0]) > 0 {
		return fmt.Sprintf("projects/%s/traces/%s", config.ProjectId, parts[0])
	}
	return ""
}

func buildHandlerWithDefaultMiddlewares(opt *Option, next http.Handler) http.Handler {
	return middleware.Logger(
		middleware.PathValidator(
			middleware.MethodValidator(
				middleware.ContentTypeValidator(
					next, opt.ContentType,
				),
				*opt.Methods,
			),
			opt.Path,
		),
	)
}

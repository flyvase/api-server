package middleware

import "net/http"

type Group struct {
	Get  http.Handler
	Post http.Handler
}

func Demux(group *Group) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			group.Get.ServeHTTP(w, r)
		case "POST":
			group.Post.ServeHTTP(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

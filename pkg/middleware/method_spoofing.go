package middleware

import "net/http"

func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			method := r.FormValue("_method")
			if method == "" {
				method = r.Header.Get("X-HTTP-Method-Override")
			}
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				r.Method = method
			}
		}
		next.ServeHTTP(w, r)
	})
}

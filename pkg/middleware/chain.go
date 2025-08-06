package middleware

import "net/http"

func Chain(h http.Handler, middlewareSlice ...func(http.Handler) http.Handler) http.Handler {
	for key := range middlewareSlice {
		h = middlewareSlice[key](h)
	}
	return h
}

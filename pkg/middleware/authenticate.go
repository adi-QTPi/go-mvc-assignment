package middleware

import (
	"fmt"
	"net/http"
)

func RequiredEntries(entries ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("the list of required entries are : ", entries)

			next.ServeHTTP(w, r)
		})
	}
}

func AnotherMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is a middleware ")

		next.ServeHTTP(w, r)
	})
}

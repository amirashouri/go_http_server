package api

import (
	"fmt"
	"net/http"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Api got called %s with %s method\n", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

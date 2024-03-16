package api

import (
	"fmt"
	"net/http"
	"strings"

	"os"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a http middleware for authorization
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			e := NewResponseError("authorization header is not provided")
			http.Error(w, e.Error(), http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			e := NewResponseError("invalid authorization header format")
			http.Error(w, e.Error(), http.StatusUnauthorized)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		accessToken := fields[1]
		fmt.Fprintln(os.Stdout, []any{"access token is: %s", accessToken}...)
		next.ServeHTTP(w, r)
	})
}

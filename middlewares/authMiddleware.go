package middlewares

import (
	"net/http"
	"strings"

	jwtValidator "mystore.com/infrastructure/security/jwtValidator"
)

func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authToken = strings.Replace(authToken, "Bearer ", "", 1)

		tokenErr := jwtValidator.Validate(authToken)
		if tokenErr != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

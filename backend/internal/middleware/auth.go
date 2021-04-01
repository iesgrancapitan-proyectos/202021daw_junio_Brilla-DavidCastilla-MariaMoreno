package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"
)

func NeedsAuth(database driver.Database, next http.HandlerFunc) http.HandlerFunc {

	keyFunc := func(t *jwt.Token) (interface{}, error) { return []byte("secret"), nil }

	return func(rw http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(rw, "Cookie not found", http.StatusUnauthorized)
			return
		}

		if cookie.Expires.After(time.Now()) {
			http.Error(rw, "Cookie expired", http.StatusUnauthorized)
			return
		}

		token := cookie.Value

		var claims jwt.StandardClaims
		if _, err := jwt.ParseWithClaims(token, &claims, keyFunc); err != nil {
			http.Error(rw, "Problem parsing JWT token", http.StatusUnauthorized)
			return
		}

		col, err := database.Collection(context.Background(), "User")
		if err != nil {
			http.Error(rw, "Error. Problem fetching collection", http.StatusUnauthorized)
			return
		}
		exists, err := col.DocumentExists(context.Background(), claims.Issuer)
		if err != nil {
			http.Error(rw, "Error. Problem fetching document", http.StatusUnauthorized)
			return
		}

		if !exists {
			http.Error(rw, "Authenticated user doesn't exists", http.StatusUnauthorized)
			return

		}

		r = r.WithContext(context.WithValue(r.Context(), "authUser", claims.Issuer))

		next.ServeHTTP(rw, r)
	}
}

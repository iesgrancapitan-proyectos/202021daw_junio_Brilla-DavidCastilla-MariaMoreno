package middleware

import (
	"context"
	"net/http"
	"os"

	// "time"

	"github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"
)

// JwtKey is the key to sign the JWT tokens
var JwtKey []byte = []byte(os.Getenv("JWT_KEY"))
var AuthUser string = "authUser"

func AuthenticatedUser(r *http.Request) string {
	return r.Context().Value(AuthUser).(string)
}

func NeedsAuth(database driver.Database, next http.HandlerFunc) http.HandlerFunc {

	keyFunc := func(t *jwt.Token) (interface{}, error) { return JwtKey, nil }

	return func(rw http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(rw, "Cookie not found", http.StatusUnauthorized)
			return
		}

		tokenRaw := cookie.Value

		var claims jwt.StandardClaims
		tkn, err := jwt.ParseWithClaims(tokenRaw, &claims, keyFunc)

		if !tkn.Valid {
			http.Error(rw, "Token not valid", http.StatusUnauthorized)
			return
		} else if err != nil && err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			// TODO: Validate refresh token, if valid generate new token
			http.Error(rw, "Token expired", http.StatusUnauthorized)
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

		r = r.WithContext(context.WithValue(r.Context(), AuthUser, claims.Issuer))

		next.ServeHTTP(rw, r)
	}
}

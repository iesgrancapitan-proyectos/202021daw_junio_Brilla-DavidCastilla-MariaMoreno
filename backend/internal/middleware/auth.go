package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	// "time"

	"github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"

	"brilla/internal/database/queries"
)

// JwtKey is the key to sign the JWT tokens
var JwtKey []byte = []byte(os.Getenv("JWT_KEY"))
var AuthUser string = "authUser"
var AuthID string = "authID"

func AuthenticatedUser(r *http.Request) (string, string) {
	return r.Context().Value(AuthUser).(string), r.Context().Value(AuthID).(string)
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
			http.Error(rw, "Token expired", http.StatusUnauthorized)
			return
		}

		cursor, err := database.Query(context.Background(), queries.GetUserQuery, map[string]interface{}{
			"username": claims.Issuer,
		})

		if driver.IsNotFound(err) {
			http.Error(rw, "Authenticated user doesn't exists", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(rw, "Error. Problem fetching document", http.StatusUnauthorized)
			return
		}

		var user map[string]interface{}
		md, err := cursor.ReadDocument(context.Background(), &user)

		fmt.Println(user)
		fmt.Println(md)

		r = r.WithContext(context.WithValue(r.Context(), AuthUser, user["username"]))
		r = r.WithContext(context.WithValue(r.Context(), AuthID, md.Key))

		next.ServeHTTP(rw, r)
	}
}

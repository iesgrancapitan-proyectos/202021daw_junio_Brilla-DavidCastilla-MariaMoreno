package router

import (
	"brilla/internal/database/queries"
	"brilla/internal/middleware"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	arango "github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var googleSecret = os.Getenv("SECRET_OAUTH")
var googleID = os.Getenv("ID_OAUTH")

var apiURL = os.Getenv("API_URL")

var googleOauthConfig = &oauth2.Config{
	ClientID:     googleID,
	ClientSecret: googleSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	RedirectURL:  apiURL + "/auth/google/confirm",
	Endpoint:     google.Endpoint,
}

type googleResponse struct {
	Email          string `json:"email"`
	Verified_email bool   `json:"verified_email"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
}

// googleAuth: /auth/google
func (server *Server) googleAuth(wr http.ResponseWriter, r *http.Request) {

	url := googleOauthConfig.AuthCodeURL("state")
	http.Redirect(wr, r, url, http.StatusTemporaryRedirect)

}

func randomString() string {
	b := make([]byte, 10)
	rand.Read(b)
	var sb strings.Builder
	base64.NewEncoder(base64.RawURLEncoding, &sb).Write(b)
	return sb.String()
}

func (server *Server) googleAuthConfirm(rw http.ResponseWriter, r *http.Request) {

	// FIX: Handle errors

	code := r.FormValue("code")
	tkn, _ := googleOauthConfig.Exchange(context.Background(), code)

	resRaw, _ := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tkn.AccessToken)
	var res googleResponse
	json.NewDecoder(resRaw.Body).Decode(&res)

	username := "estrella-" + randomString()
	// validación email
	email := res.Email
	bio := ""
	name := res.Name

	profileImg := res.Picture

	_, err := server.database.Query(context.Background(), queries.InsertUserQuery, map[string]interface{}{
		"username":    username,
		"email":       email,
		"bio":         bio,
		"password":    nil,
		"name":        name,
		"birthday":    nil,
		"profile_img": profileImg,
		"is_active":   true,
	})
	if arango.IsConflict(err) {
		writeError(rw, "Error. Conflict user. "+err.Error(), http.StatusConflict)
		return
	} else if err != nil {
		writeError(rw, "Error. Creating user. "+err.Error(), http.StatusConflict)
		return
	}

	expirationTime := time.Now().AddDate(0, 0, 3)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: expirationTime.Unix(),
	})

	signed, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		writeError(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "token",
		Value:    signed,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expirationTime,
		Path:     "/",
	})

	rw.Header().Add("X-Token", signed)

	http.Redirect(rw, r, "/", http.StatusFound)
	json.NewEncoder(rw).Encode(map[string]string{
		"username": username,
	})

}

// FACEBOOK
var facebookSecret = os.Getenv("SECRET_FACEBOOK")
var facebookID = os.Getenv("ID_FACEBOOK")

var facebookOauthConfig = &oauth2.Config{
	ClientID:     facebookID,
	ClientSecret: facebookSecret,
	Scopes:       []string{"email"},
	RedirectURL:  apiURL + "/auth/facebook/confirm",
	Endpoint:     facebook.Endpoint,
}

//facebookAuth: /auth/facebook
func (server *Server) facebookAuth(wr http.ResponseWriter, r *http.Request) {

	url := facebookOauthConfig.AuthCodeURL("state")
	http.Redirect(wr, r, url, http.StatusTemporaryRedirect)

}

func (server *Server) facebookAuthConfirm(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	tkn, _ := facebookOauthConfig.Exchange(context.Background(), r.FormValue("code"))

	resRaw, _ := http.Get("https://graph.facebook.com/me?access_token=" + tkn.AccessToken)
	x, _ := json.Marshal(resRaw.Body)
	fmt.Printf("%#v\n", string(x))
	// var res googleResponse
	// json.NewDecoder(resRaw.Body).Decode(&res)

	// username := "estrella-" + randomString()
	// // validación email
	// email := res.Email
	// bio := ""
	// name := res.Name

	// profileImg := res.Picture

	// _, err := server.database.Query(context.Background(), queries.InsertUserQuery, map[string]interface{}{
	// 	"username":    username,
	// 	"email":       email,
	// 	"bio":         bio,
	// 	"password":    nil,
	// 	"name":        name,
	// 	"birthday":    nil,
	// 	"profile_img": profileImg,
	// 	"is_active":   true,
	// })
	// if arango.IsConflict(err) {
	// 	writeError(rw, "Error. Conflict user. "+err.Error(), http.StatusConflict)
	// 	return
	// } else if err != nil {
	// 	writeError(rw, "Error. Creating user. "+err.Error(), http.StatusConflict)
	// 	return
	// }

	// expirationTime := time.Now().AddDate(0, 0, 3)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    username,
	// 	ExpiresAt: expirationTime.Unix(),
	// })

	// signed, err := token.SignedString(middleware.JwtKey)
	// if err != nil {
	// 	writeError(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	// }

	// http.SetCookie(rw, &http.Cookie{
	// 	Name:     "token",
	// 	Value:    signed,
	// 	Secure:   false,
	// 	HttpOnly: true,
	// 	SameSite: http.SameSiteLaxMode,
	// 	Expires:  expirationTime,
	// 	Path:     "/",
	// })

	// rw.Header().Add("X-Token", signed)

	// http.Redirect(rw, r, "/", http.StatusFound)
	// json.NewEncoder(rw).Encode(map[string]string{
	// 	"username": username,
	// })

}

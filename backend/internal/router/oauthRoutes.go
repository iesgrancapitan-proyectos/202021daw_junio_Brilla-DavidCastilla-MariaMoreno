package router

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleID = os.Getenv("SECRET_OAUTH")
var googleSecret = os.Getenv("ID_OAUTH")

var apiURL = os.Getenv("API_URL")

// googleAuth: /auth/google
func (server *Server) googleAuth(wr http.ResponseWriter, r *http.Request) {

	conf := &oauth2.Config{
		ClientID:     googleID,
		ClientSecret: googleSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, //, "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  apiURL + "/auth/google/confirm",
		Endpoint:     google.Endpoint,
	}

	url := conf.AuthCodeURL("state")
	http.Redirect(wr, r, url, http.StatusTemporaryRedirect)

}

func (server *Server) googleAuthConfirm(wr http.ResponseWriter, r *http.Request) {

	println("asdf")

}

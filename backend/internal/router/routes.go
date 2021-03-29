package router

import "net/http"

func getUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return info of user in JSON
}

func getUserBrights(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return brights of user in JSON
}

func postUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Upload a new user
}

func postUserFollow(rw http.ResponseWriter, r *http.Request) {
	// TODO: Follows a user
}

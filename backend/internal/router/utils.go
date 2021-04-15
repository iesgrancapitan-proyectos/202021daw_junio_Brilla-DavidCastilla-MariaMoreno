package router

import (
	"encoding/json"
	"net/http"
)

// loginUserBody for route postLogin
type loginUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// activateUserBody for route postActivateUser
type activateUserBody struct {
	Token string `json:"token"`
}

func writeError(rw http.ResponseWriter, err string, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(map[string]string{"error": err})
}

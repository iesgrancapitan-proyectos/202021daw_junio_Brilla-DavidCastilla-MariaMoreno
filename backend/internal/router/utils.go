package router

import (
	"encoding/json"
	"net/http"
)

func writeError(rw http.ResponseWriter, err string, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(map[string]string{"error": err})
}

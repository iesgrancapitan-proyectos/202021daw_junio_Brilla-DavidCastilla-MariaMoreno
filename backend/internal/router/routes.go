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
func postRebrilla(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crea un brillo pasado por JSON
}

func postInteraction(rw http.ResponseWriter, r *http.Request) {
	//TODO : Interactua con un brillo y el tipo de intereción por el json mandado
}

func postComment(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crear un brillo respuesta a otro Brillo
}

func postBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Crea un brillo
}

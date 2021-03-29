package router

import "net/http"

func (server *Server) getBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return info of user in JSON
}

func (server *Server) getUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return info of user in JSON
}

func (server *Server) getUserBrights(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return brights of user in JSON
}

func (server *Server) postUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Upload a new user
}

func (server *Server) postLogin(rw http.ResponseWriter, r *http.Request) {
	// TODO: Handle login
}

func (server *Server) postUserFollow(rw http.ResponseWriter, r *http.Request) {
	// TODO: Follows a user
}
func (server *Server) postRebrilla(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crea un brillo pasado por JSON
}

func (server *Server) postInteraction(rw http.ResponseWriter, r *http.Request) {
	//TODO : Interactua con un brillo y el tipo de intereción por el json mandado
}

func (server *Server) postComment(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crear un brillo respuesta a otro Brillo
}

func (server *Server) postBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Crea un brillo
}

func (server *Server) deleteUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
}

func (server *Server) deleteBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
}

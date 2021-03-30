package router

import (
	"context"
	"net/http"

	"github.com/matthewhartstonge/argon2"

	"brilla/internal/database/queries"
	"brilla/internal/models"
)

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
	
}

func (server *Server) postLogin(rw http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", 500)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	cursor, err := server.database.Query(context.Background(), queries.GetUserQuery, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		http.Error(rw, "Error fetching database", 500)
		return
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		http.Error(rw, "Error user not found", 404)
		return
	}

	var user models.User
	cursor.ReadDocument(context.Background(), &user)

	match, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil {
		http.Error(rw, "Error verifyng password", 500)
		return
	}

	if !match {
		http.Error(rw, "Error: Incorrect password", 401)
		return
	}

	// TODO: Generate token

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

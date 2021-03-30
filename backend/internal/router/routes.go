package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/matthewhartstonge/argon2"

	"brilla/internal/database/queries"
	"brilla/internal/models"
)

//getBright route: /brights
func (server *Server) getBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return info of user in JSON

}

// getUser route: /user/:username
func (server *Server) getUser(rw http.ResponseWriter, r *http.Request) {
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	collection, err := server.database.Collection(context.Background(), "User")
	if err != nil {
		http.Error(rw, "Error can not find collection", 500)
		return
	}

	var user models.User
	_, err = collection.ReadDocument(context.Background(), username, &user)
	if err != nil {
		http.Error(rw, "Error can not read collection", 500)
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		http.Error(rw, "Error encoding json", 500)
		return
	}
}

// getUserBrights route: /user/:username/bright
func (server *Server) getUserBrights(rw http.ResponseWriter, r *http.Request) {
	// TODO: Return brights of user in JSON
}

// postUser route: /user
func (server *Server) postUser(rw http.ResponseWriter, r *http.Request) {

}

// postLogin route: /user/login
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

	_, err = argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil {
		http.Error(rw, "Error: Incorrect password", 401)
		return
	}

	// TODO: Generate token

}

//putUserFollor route: /user/:username/follow
func (server *Server) postUserFollow(rw http.ResponseWriter, r *http.Request) {
	// TODO: Follows a user
}

//postRebrilla route: /brights/rebrilla
func (server *Server) postRebrilla(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crea un brillo pasado por JSON
}

// postInteraction route: /brights/interaction
func (server *Server) postInteraction(rw http.ResponseWriter, r *http.Request) {
	//TODO : Interactua con un brillo y el tipo de intereción por el json mandado
}

// postComment route: /brights/comment
func (server *Server) postComment(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crear un brillo respuesta a otro Brillo
}

//postBright route: /brights
func (server *Server) postBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Crea un brillo
}

//deleteUser route: /user/delete
func (server *Server) deleteUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
}

//deleteBright route: /brights/delete
func (server *Server) deleteBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
}

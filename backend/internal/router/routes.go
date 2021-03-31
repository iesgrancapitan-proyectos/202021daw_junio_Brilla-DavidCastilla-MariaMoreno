package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	arango "github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/matthewhartstonge/argon2"

	"brilla/internal/database/queries"
	"brilla/internal/models"
)

//getBright route: /brights/:id
func (server *Server) getBright(rw http.ResponseWriter, r *http.Request) {
	idBrillo := httprouter.ParamsFromContext(r.Context()).ByName("id")

	collection, err := server.database.Collection(context.Background(), "Brillo")
	if err != nil {
		http.Error(rw, "Error can not find collection", http.StatusInternalServerError)
		return
	}

	var brillo models.Brillo

	if _, err = collection.ReadDocument(context.Background(), idBrillo, &brillo); arango.IsNotFound(err) {
		http.Error(rw, "Error: Bright not found. "+err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Error can not read collection. "+err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(brillo)
	if err != nil {
		http.Error(rw, "Error encoding json", http.StatusInternalServerError)
		return
	}

}

// getUser route: /user/:username
func (server *Server) getUser(rw http.ResponseWriter, r *http.Request) {
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	collection, err := server.database.Collection(context.Background(), "User")
	if err != nil {
		http.Error(rw, "Error can not find collection", http.StatusInternalServerError)
		return
	}

	var user models.User
	_, err = collection.ReadDocument(context.Background(), username, &user)
	if arango.IsNotFound(err) {
		http.Error(rw, "Error: User not found. "+err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Error can not read collection", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		rw.Header()
		http.Error(rw, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

// getUserBrights route: /user/:username/brights
func (server *Server) getUserBrights(rw http.ResponseWriter, r *http.Request) {
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")
	cursor, err := server.database.Query(context.Background(), queries.GetBrillosByAuthorQuery, map[string]interface{}{"username": username})
	if err != nil {
		http.Error(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

	brillos := make([]models.Brillo, 0)
	for cursor.HasMore() {
		var brillo models.Brillo
		cursor.ReadDocument(context.Background(), &brillo)
		brillos = append(brillos, brillo)
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(brillos)
	if err != nil {
		rw.Header()
		http.Error(rw, "Error encoding json", http.StatusInternalServerError)
		return
	}

}

// postUser route: /user
func (server *Server) postUser(rw http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	bio := r.FormValue("bio")
	password := r.FormValue("password")
	argon := argon2.DefaultConfig()
	password_hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		http.Error(rw, "Error can not hash password", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	birthday, err := strconv.Atoi(r.FormValue("birthday"))
	if err != nil {
		http.Error(rw, "Error date is not a number", http.StatusBadRequest)
		return
	}

	profileImg := r.FormValue("profileImg")

	collection, err := server.database.Collection(context.Background(), "User")
	if err != nil {
		http.Error(rw, "Error can not find collection", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Username:   username,
		Email:      email,
		Password:   string(password_hash),
		Name:       name,
		Bio:        bio,
		Birthday:   int64(birthday),
		ProfileImg: profileImg,
	}

	collection.CreateDocument(r.Context(), &user)
	fmt.Fprint(rw, "success")
}

// postLogin route: /user/login
func (server *Server) postLogin(rw http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	collection, err := server.database.Collection(context.Background(), "User")
	if err != nil {
		http.Error(rw, "Error can not find collection", http.StatusInternalServerError)
		return
	}

	var user models.User

	if _, err = collection.ReadDocument(context.Background(), username, &user); arango.IsNotFound(err) {
		http.Error(rw, "Error: User not found. "+err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Error can not read collection. "+err.Error(), http.StatusInternalServerError)
		return
	}

	match, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil {
		http.Error(rw, "Error: Incorrect password", http.StatusUnauthorized)
		return
	}

	if !match {
		http.Error(rw, "Error: Incorrect password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: time.Now().AddDate(0, 0, 3).Unix(),
	})

	signed, err := token.SignedString([]byte("secret"))
	if err != nil {
		http.Error(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "token",
		Value:    signed,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().AddDate(0, 0, 3),
	})

	fmt.Fprint(rw, "success")

}

//putUserFollor route: /user/:username/follow
func (server *Server) putUserFollow(rw http.ResponseWriter, r *http.Request) {
	// TODO: Follows a user
	follower := r.Context().Value("authUser").(string)

	followed := httprouter.ParamsFromContext(r.Context()).ByName("username")

	vars := map[string]interface{}{
		"follower": follower,
		"followed": followed,
	}

	cursor, err := server.database.Query(arango.WithQueryCount(context.Background()), queries.IsFollowingQuery, vars)

	if cursor.Count() >= 1 {
		//si ya lo sigue unfoollowed
		return
	}

	_, err = server.database.Query(context.Background(), queries.NewFollowQuery, vars)
	if err != nil {
		http.Error(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

}

//postRebrilla route: /brights/rebrilla
func (server *Server) postRebrilla(rw http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("authUser").(string)

	//obtener el id del brillo pulsado
	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	brilloId := r.FormValue("brilloId")

	_, err = server.database.Query(context.Background(), queries.RebrilloQuery, map[string]interface{}{
		"userId":   username,
		"brilloId": brilloId,
	})
	if err != nil {
		http.Error(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

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
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	

}

//deleteUser route: /user/delete
func (server *Server) deleteUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright

}

//deleteBright route: /brights/delete
func (server *Server) deleteBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
}

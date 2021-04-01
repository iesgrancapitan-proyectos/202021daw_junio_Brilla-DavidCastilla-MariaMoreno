package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	cursor, err := server.database.Query(context.Background(), queries.GetUserQuery, map[string]interface{}{"username": username})
	if err != nil {
		http.Error(rw, "Error can not read collection", http.StatusInternalServerError)
		return
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		http.Error(rw, "Error: User not found.", http.StatusNotFound)
		return
	}

	// FIX: Return to this later
	var user map[string]interface{}
	cursor.ReadDocument(context.Background(), &user)

	rw.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(rw).Encode(user); err != nil {
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
	defer cursor.Close()

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
		Path:     "/",
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
		//si ya lo sigue unfollowed
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

	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	brilloId := r.FormValue("brilloId")

	_, err = server.database.Query(context.Background(), queries.RebrilloQuery, map[string]interface{}{
		"username":  username,
		"brilloKey": brilloId,
	})
	if err != nil {
		http.Error(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

}

// postInteraction route: /brights/interaction
func (server *Server) postInteraction(rw http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("authUser").(string)

	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	brilloKey := r.FormValue("brilloKey")
	interaction := r.FormValue("interaction")

	_, err = server.database.Query(context.Background(), queries.InteractionQuery, map[string]interface{}{
		"username":  username,
		"brilloKey": brilloKey,
		"type":      interaction,
	})
	if err != nil {
		http.Error(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

}

//postBright route: /brights
func (server *Server) postBright(rw http.ResponseWriter, r *http.Request) {

	username := r.Context().Value("authUser").(string)

	err := r.ParseMultipartForm(8 >> 20)
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	headers := r.MultipartForm.File["media"]
	if len(headers) > 4 {
		http.Error(rw, "Cannot upload more than 4 files", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	media := make([]string, 0, 4)

	for i, h := range headers {
		srcFile, _ := h.Open()
		dstFilepath := "/media/" + username + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + strconv.Itoa(i)
		dstFile, _ := os.OpenFile(dstFilepath, os.O_CREATE|os.O_WRONLY, 0755)
		io.Copy(dstFile, srcFile)

		srcFile.Close()
		dstFile.Close()

		media = append(media, dstFilepath)
	}

	_, err = server.database.Query(context.Background(), queries.NewBrilloQuery, map[string]interface{}{
		"username": username,
		"content":  content,
		"media":    media,
	})
	if err != nil {
		http.Error(rw, "Error inserting to database", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(rw, "success")

}

// postComment route: /brights/comment
func (server *Server) postComment(rw http.ResponseWriter, r *http.Request) {
	//TODO: Crear un brillo respuesta a otro Brillo
	username := r.Context().Value("username").(string)

	err := r.ParseMultipartForm(8 >> 20)
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	headers := r.MultipartForm.File["media"]
	if len(headers) > 4 {
		http.Error(rw, "Cannot upload more than 4 files", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	media := make([]string, 0, 4)

	for i, h := range headers {
		srcFile, _ := h.Open()
		dstFilepath := "/media/" + username + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + strconv.Itoa(i)
		dstFile, _ := os.OpenFile(dstFilepath, os.O_CREATE|os.O_WRONLY, 0755)
		io.Copy(dstFile, srcFile)

		srcFile.Close()
		dstFile.Close()

		media = append(media, dstFilepath)
	}

	//brillo al que responde
	brilloKey := r.FormValue("brilloKey")

	_, err = server.database.Query(context.Background(), queries.CommentBrilloQuery, map[string]interface{}{
		"username":  username,
		"content":   content,
		"media":     media,
		"brilloKey": brilloKey,
	})
	if err != nil {
		http.Error(rw, "Error inserting to database", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(rw, "success")

}

//deleteUser route: /user/delete
func (server *Server) deleteUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
	username := r.Context().Value("username").(string)

	_, err := server.database.Query(context.Background(), queries.DeactivateUserQuery, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		http.Error(rw, "Error inserting to database", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(rw, "success")
}

//deleteBright route: /brights/delete
func (server *Server) deleteBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright

	err := r.ParseForm()
	if err != nil {
		http.Error(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	idbrillo := r.FormValue("idbrillo")

	_, err = server.database.Query(context.Background(), queries.DeleteBrilloQuery, map[string]interface{}{
		"brilloKey": idbrillo,
	})

	if err != nil {
		http.Error(rw, "Error inserting to database", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(rw, "success")
}

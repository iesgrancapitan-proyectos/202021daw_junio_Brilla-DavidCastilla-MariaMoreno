package router

import (
	"brilla/internal/database/queries"
	"brilla/internal/middleware"
	"brilla/internal/models"
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
)

// getUser route: /user/:username
func (server *Server) getUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	cursor, err := server.database.Query(arango.WithKeepNull(context.Background(), true), queries.GetUserQuery, map[string]interface{}{"username": username})
	if err != nil {
		writeError(rw, "Can't read collection", http.StatusInternalServerError)
		return
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		writeError(rw, "User not found", http.StatusNotFound)
		return
	}

	// FIX: Return to this later
	var user map[string]interface{}
	cursor.ReadDocument(context.Background(), &user)

	delete(user, "password")

	if err = json.NewEncoder(rw).Encode(user); err != nil {
		writeError(rw, "Can't encode JSON", http.StatusInternalServerError)
		return
	}
}

// getUserBrights route: /user/:username/brights
func (server *Server) getUserBrights(rw http.ResponseWriter, r *http.Request) {
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	cursor, err := server.database.Query(context.Background(), queries.GetBrillosByAuthorQuery, map[string]interface{}{
		"username": username,
		"offset":   offset,
		"limit":    limit,
	})

	if err != nil {
		writeError(rw, "Can not connect with database"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close()

	brights := make([]map[string]interface{}, 0, 10)
	for cursor.HasMore() {
		var brillo map[string]interface{}
		cursor.ReadDocument(context.Background(), &brillo)
		brights = append(brights, brillo)
	}
	fmt.Println(brights)
	json.NewEncoder(rw).Encode(brights)

}

// postUser route: /user
func (server *Server) postUser(rw http.ResponseWriter, r *http.Request) {
	// FIX: Validate inputs

	err := r.ParseForm()
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	// validación email
	email := r.FormValue("email")
	bio := r.FormValue("bio")
	password := r.FormValue("password")
	argon := argon2.DefaultConfig()
	password_hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		writeError(rw, "Error can not hash password", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	birthday, err := strconv.Atoi(r.FormValue("birthday"))
	if err != nil {
		writeError(rw, "Error date is not a number", http.StatusBadRequest)
		return
	}

	profileImg := r.FormValue("profileImg")

	_, err = server.database.Query(context.Background(), queries.InsertUserQuery, map[string]interface{}{
		"username":    username,
		"email":       email,
		"bio":         bio,
		"password":    string(password_hash),
		"name":        name,
		"birthday":    birthday,
		"profile_img": profileImg,
	})
	if arango.IsConflict(err) {
		writeError(rw, "Error. Creating user. "+err.Error(), http.StatusConflict)
		return
	} else if err != nil {
		writeError(rw, "Error. Creating user. "+err.Error(), http.StatusConflict)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{Issuer: username})
	signed, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		writeError(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	}

	err = sendMail(email, signed)
	if err != nil {
		writeError(rw, "Error sending email. "+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// postUser route: /user/edit
func (server *Server) postUserEdit(rw http.ResponseWriter, r *http.Request) {
	// FIX: Validate inputs
	_, key := middleware.AuthenticatedUser(r)

	err := r.ParseMultipartForm(4 >> 20)
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	println(username)

	bio := r.FormValue("bio")
	// password := r.FormValue("password")
	// argon := argon2.DefaultConfig()
	// password_hash, err := argon.HashEncoded([]byte(password))
	// if err != nil {
	// 	writeError(rw, "Error can not hash password", http.StatusInternalServerError)
	// 	return
	// }

	name := r.FormValue("name")
	// birthday, err := strconv.Atoi(r.FormValue("birthday"))
	// if err != nil {
	// 	writeError(rw, "Error date is not a number", http.StatusBadRequest)
	// 	return
	// }

	// profileImg := r.FormValue("profileImg")

	print("LLAVE: ")
	println(key)

	_, err = server.database.Query(context.Background(), queries.UpdateUserQuery, map[string]interface{}{
		"username": key,
		"newUser":  username,
		"bio":      bio,
		// "password":    string(password_hash),
		"name": name,
		// "birthday":    birthday,
		// "profile_img": profileImg,
	})
	if arango.IsConflict(err) {
		writeError(rw, "Error. Conflict updating user. "+err.Error(), http.StatusConflict)
		return
	} else if err != nil {
		writeError(rw, "Error. Updating user. "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, h := range r.MultipartForm.File["profile"] {

		file, err := h.Open()
		if err != nil {
			// TODO: Manage err
		}

		dirPath := "/media/" + key + "/"
		os.MkdirAll(dirPath, 0655)
		dstFilepath := dirPath + "pp"
		println(dstFilepath)
		dstFile, err := os.OpenFile(dstFilepath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			// TODO: Manage err
		}
		io.Copy(dstFile, file)

		file.Close()
		dstFile.Close()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{Issuer: username})
	signed, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		writeError(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	}

	expirationTime := time.Now().AddDate(0, 0, 3)

	http.SetCookie(rw, &http.Cookie{
		Name:     "token",
		Value:    signed,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expirationTime,
		Path:     "/",
	})

	rw.Header().Add("X-Token", signed)

	rw.WriteHeader(http.StatusOK)
}

// postActivateUser router: /user/activate
func (server *Server) postActivateUser(rw http.ResponseWriter, r *http.Request) {

	var activateUserBody activateUserBody
	err := json.NewDecoder(r.Body).Decode(&activateUserBody)
	if err != nil {
		writeError(rw, "Can't parse JSON. "+err.Error(), http.StatusInternalServerError)
		return
	}

	var claims jwt.StandardClaims
	tkn, err := jwt.ParseWithClaims(activateUserBody.Token, &claims, func(t *jwt.Token) (interface{}, error) { return middleware.JwtKey, nil }) // FIX: New jwtKey for activation
	if !tkn.Valid {
		writeError(rw, "Token not valid", http.StatusBadRequest)
		return
	} else if err != nil && err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
		writeError(rw, "Token expired", http.StatusRequestTimeout)
		return
	} else if err != nil {
		writeError(rw, "Unknown error. "+err.Error(), http.StatusInternalServerError)
		return
	}

	username := claims.Issuer

	_, err = server.database.Query(context.Background(), queries.ActivateUserQuery, map[string]interface{}{"username": username})
	if err != nil {
		writeError(rw, "Error activating user", http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().AddDate(0, 0, 3)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: expirationTime.Unix(),
	})

	signed, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		writeError(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "token",
		Value:    signed,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expirationTime,
		Path:     "/",
	})

	rw.Header().Add("X-Token", signed)

	json.NewEncoder(rw).Encode(map[string]string{
		"username": username,
	})

}

// postLogin route: /user/login
func (server *Server) postLogin(rw http.ResponseWriter, r *http.Request) {

	var postLoginBody loginUserBody
	err := json.NewDecoder(r.Body).Decode(&postLoginBody)
	if err != nil {
		writeError(rw, "Problem parsing JSON. "+err.Error(), http.StatusInternalServerError)
		return
	}

	username := postLoginBody.Username

	var user map[string]interface{}

	cursor, err := server.database.Query(arango.WithQueryCount(context.Background()), queries.GetUserQuery, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		writeError(rw, "Error can not read collection. "+err.Error(), http.StatusInternalServerError)
		return
	}

	if cursor.Count() == 0 {
		writeError(rw, "Error: User not found.", http.StatusNotFound)
		return
	}

	md, err := cursor.ReadDocument(context.Background(), &user)
	id := md.ID.String()

	match, err := argon2.VerifyEncoded([]byte(postLoginBody.Password), []byte(user["password"].(string)))
	if err != nil || !match {
		writeError(rw, "Error: Incorrect password. ", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().AddDate(0, 0, 3)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    username,
		ExpiresAt: expirationTime.Unix(),
	})

	signed, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		writeError(rw, "Error signing token. "+err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "token",
		Value:    signed,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expirationTime,
		Path:     "/",
	})

	rw.Header().Add("X-Token", signed)

	json.NewEncoder(rw).Encode(map[string]string{
		"username": username,
		"id":       id,
	})

}

//para seguir o dejar de seguir cuando pulse el boton
//putUserFollor route: /user/:username/follow
func (server *Server) putUserFollow(rw http.ResponseWriter, r *http.Request) {
	// TODO: Follows a user
	_, follower := middleware.AuthenticatedUser(r)

	followed := httprouter.ParamsFromContext(r.Context()).ByName("username")

	vars := map[string]interface{}{
		"follower": follower,
		"followed": followed,
	}

	_, err := server.database.Query(context.Background(), queries.NewFollowQuery, vars)
	if arango.IsConflict(err) {
		_, err = server.database.Query(context.Background(), queries.DeleteFollowQuery, vars)
		if err != nil {
			writeError(rw, "Error can not connect with database. "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]bool{
			"result": false,
		})
		return
	}

	json.NewEncoder(rw).Encode(map[string]bool{
		"result": true,
	})

}

//para saber si se siguen o no para lo q pone en el boton
//isFollowing route: /user/:username/isFolllowing
func (server *Server) isFollowing(rw http.ResponseWriter, r *http.Request) {
	// TODO: Follows a user
	_, follower := middleware.AuthenticatedUser(r)

	followed := httprouter.ParamsFromContext(r.Context()).ByName("username")

	vars := map[string]interface{}{
		"follower": follower,
		"followed": followed,
	}
	cursor, err := server.database.Query(arango.WithQueryCount(context.Background()), queries.IsFollowingQuery, vars)

	if err != nil {
		writeError(rw, "Error in query. "+err.Error(), http.StatusInternalServerError)
		return
	}

	result := false
	if cursor.Count() >= 1 {
		//si ya lo sigue unfollowed
		result = true
	}
	json.NewEncoder(rw).Encode(map[string]bool{
		"follow": result,
	})

}

//deleteUser route: /user/delete
func (server *Server) deleteUser(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright
	_, username := middleware.AuthenticatedUser(r)

	_, err := server.database.Query(context.Background(), queries.DeactivateUserQuery, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		writeError(rw, "Error delete user from database ", http.StatusInternalServerError)
		return
	}

}

//getRefresh /refresh
func (_ *Server) getRefresh(rw http.ResponseWriter, r *http.Request) {
	username, key := middleware.AuthenticatedUser(r)
	json.NewEncoder(rw).Encode(map[string]string{"username": username, "key": key})
}

//getLogout /logout
func (_ *Server) getLogout(rw http.ResponseWriter, r *http.Request) {
	http.SetCookie(rw, &http.Cookie{
		Name:     "token",
		Value:    "",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Path:     "/",
	})
}

// /nfollowers/:username
func (server *Server) getFollowers(rw http.ResponseWriter, r *http.Request) {
	//username := middleware.AuthenticatedUser(r)
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	cursor, err := server.database.Query(context.Background(), queries.NFollowers, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		writeError(rw, "Error get followers from database ", http.StatusInternalServerError)
		return
	}

	var result int

	_, err = cursor.ReadDocument(context.Background(), &result)
	if err != nil {
		writeError(rw, "Error reading document ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(map[string]int{
		"followers": result,
	})

}

// /nfollowed/:username
func (server *Server) getFollowed(rw http.ResponseWriter, r *http.Request) {
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	cursor, err := server.database.Query(context.Background(), queries.NFollowed, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		writeError(rw, "Error delete user from database ", http.StatusInternalServerError)
		return
	}

	var result int

	_, err = cursor.ReadDocument(context.Background(), &result)
	if err != nil {
		writeError(rw, "Error reading document ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(map[string]int{
		"followed": result,
	})

}

// /user/:username/brights/count
func (server *Server) getNumBrillos(rw http.ResponseWriter, r *http.Request) {
	username := httprouter.ParamsFromContext(r.Context()).ByName("username")

	cursor, err := server.database.Query(arango.WithQueryCount(context.Background(), true), queries.GetBrillosByAuthorQuery, map[string]interface{}{"username": username, "offset": 0, "limit": 4294967295})
	if err != nil {
		writeError(rw, "Can not connect with database"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close()

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(map[string]int64{
		"nbrillos": cursor.Count(),
	})
	if err != nil {
		writeError(rw, "Encoding JSON", http.StatusInternalServerError)
		return
	}

}

//email exits route: /user/exits
func (server *Server) getEmailExits(rw http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeError(rw, "Problem parsing JSON. "+err.Error(), http.StatusInternalServerError)
		return
	}

	email := user.Email
	cursor, err := server.database.Query(arango.WithQueryCount(context.Background(), true), queries.ReturnEmailExits, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		writeError(rw, "Error get emails from database ", http.StatusInternalServerError)
		return
	}

	if cursor.Count() == 1 {
		writeError(rw, "Error, eamil alrady exits ", http.StatusConflict)
	}

}

package router

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	arango "github.com/arangodb/go-driver"
	"github.com/julienschmidt/httprouter"

	"brilla/internal/database/queries"
	"brilla/internal/middleware"
	"brilla/internal/models"
)

//getBright route: /brights/:id
func (server *Server) getBright(rw http.ResponseWriter, r *http.Request) {
	idBrillo := httprouter.ParamsFromContext(r.Context()).ByName("id")

	collection, err := server.database.Collection(context.Background(), "Brillo")
	if err != nil {
		writeError(rw, "Error can not find collection", http.StatusInternalServerError)
		return
	}

	var brillo models.Brillo

	if _, err = collection.ReadDocument(context.Background(), idBrillo, &brillo); arango.IsNotFound(err) {
		writeError(rw, "Error: Bright not found. "+err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		writeError(rw, "Error can not read collection. "+err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(brillo)
	if err != nil {
		writeError(rw, "Error encoding json", http.StatusInternalServerError)
		return
	}

}

//postRebrilla route: /brights/rebrilla
func (server *Server) postRebrilla(rw http.ResponseWriter, r *http.Request) {

	username := middleware.AuthenticatedUser(r)

	err := r.ParseForm()
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	brilloId := r.FormValue("brilloId")

	_, err = server.database.Query(context.Background(), queries.RebrilloQuery, map[string]interface{}{
		"username":  username,
		"brilloKey": brilloId,
	})
	if err != nil {
		writeError(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

}

// postInteraction route: /brights/interaction
func (server *Server) postInteraction(rw http.ResponseWriter, r *http.Request) {
	username := middleware.AuthenticatedUser(r)

	err := r.ParseForm()
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
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
		writeError(rw, "Error can not connect with database", http.StatusInternalServerError)
		return
	}

}

//postBright route: /brights
func (server *Server) postBright(rw http.ResponseWriter, r *http.Request) {

	username := middleware.AuthenticatedUser(r)

	err := r.ParseMultipartForm(8 >> 20)
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	headers := r.MultipartForm.File["media"]
	if len(headers) > 4 {
		writeError(rw, "Cannot upload more than 4 files", http.StatusBadRequest)
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
		writeError(rw, "Error inserting to database", http.StatusInternalServerError)
		return
	}

}

// postComment route: /brights/comment
func (server *Server) postComment(rw http.ResponseWriter, r *http.Request) {
	username := middleware.AuthenticatedUser(r)

	err := r.ParseMultipartForm(8 >> 20)
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	headers := r.MultipartForm.File["media"]
	if len(headers) > 4 {
		writeError(rw, "Cannot upload more than 4 files", http.StatusBadRequest)
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
		writeError(rw, "Error inserting to database", http.StatusInternalServerError)
		return
	}

}

//deleteBright route: /brights/:idbrillo/delete
func (server *Server) deleteBright(rw http.ResponseWriter, r *http.Request) {
	// TODO: Remove bright

	err := r.ParseForm()
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	idbrillo := httprouter.ParamsFromContext(r.Context()).ByName("idbrillo")

	_, err = server.database.Query(context.Background(), queries.DeleteBrilloQuery, map[string]interface{}{
		"brilloKey": idbrillo,
	})
	if err != nil {
		writeError(rw, "Delete brillo from database. "+err.Error(), http.StatusInternalServerError)
		return
	}

}

// getTimeline router: /timeline
func (server *Server) getTimeline(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	username := middleware.AuthenticatedUser(r)

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	cursor, err := server.database.Query(context.Background(), queries.GetTimelineQuery, map[string]interface{}{
		"username": username,
		"limit":    limit,
		"offset":   offset,
	})
	if err != nil {
		writeError(rw, "Can't query database. "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !cursor.HasMore() {
		rw.Write([]byte("[]"))
		return
	}

	brights := make([]map[string]interface{}, 0, 10)

	for cursor.HasMore() {
		var u map[string]interface{}
		_, err := cursor.ReadDocument(context.Background(), &u)
		if err != nil {
			continue
		}

		brights = append(brights, u)
	}

	json.NewEncoder(rw).Encode(brights)

}
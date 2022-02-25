package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	arango "github.com/arangodb/go-driver"
	"github.com/julienschmidt/httprouter"

	"brilla/internal/database/queries"
	"brilla/internal/middleware"
)

//getBright route: /brights/:id
func (server *Server) getBright(rw http.ResponseWriter, r *http.Request) {

	idBrillo := httprouter.ParamsFromContext(r.Context()).ByName("id")

	collection, err := server.database.Query(context.Background(), queries.GetBrilloByIdQuery, map[string]interface{}{
		"id": idBrillo,
	})
	if err != nil {
		writeError(rw, "Error can not find collection "+err.Error(), http.StatusInternalServerError)
		return
	}

	var brillo map[string]interface{}

	if _, err = collection.ReadDocument(context.Background(), &brillo); arango.IsNotFound(err) {
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

	_, username := middleware.AuthenticatedUser(r)

	err := r.ParseForm()
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	brilloId := r.FormValue("brilloId")
	println(brilloId)
	_, err = server.database.Query(context.Background(), queries.RebrilloQuery, map[string]interface{}{
		"username":  username,
		"brilloKey": brilloId,
	})

	rw.Header().Add("Content-Type", "application/json")
	if arango.IsConflict(err) {
		_, err = server.database.Query(context.Background(), queries.DeleteRebrilloQuery, map[string]interface{}{
			"username":  username,
			"brilloKey": brilloId,
		})
		if err != nil {
			writeError(rw, "Error can not connect with database. "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]bool{
			"inserted": false,
		})
		return
	} else if err != nil {
		writeError(rw, "Error can not connect with database. "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(map[string]bool{
		"inserted": true,
	})

}

// postInteraction route: /brights/interaction
func (server *Server) postInteraction(rw http.ResponseWriter, r *http.Request) {
	_, username := middleware.AuthenticatedUser(r)

	err := r.ParseForm()
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	brilloKey := r.FormValue("brilloId")
	interaction := r.FormValue("type")

	_, err = server.database.Query(context.Background(), queries.InteractionQuery, map[string]interface{}{
		"username":  username,
		"brilloKey": brilloKey,
		"type":      interaction,
	})

	rw.Header().Add("Content-Type", "application/json")

	if arango.IsConflict(err) {
		cursor, err := server.database.Query(context.Background(), `FOR i IN Interactions FILTER i._from == CONCAT('User/', @username) && i._to == CONCAT('Brillo/', @brilloKey) RETURN i`, map[string]interface{}{
			"username":  username,
			"brilloKey": brilloKey,
			// "type":      interaction,
		})
		if err != nil {
			writeError(rw, "Error can not connect with database. "+err.Error(), http.StatusInternalServerError)
			return
		}

		var readedInteraction map[string]interface{}
		metadata, err := cursor.ReadDocument(context.Background(), &readedInteraction)
		if err != nil {
			writeError(rw, "Error reading cursor. "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(readedInteraction)
		if interaction == readedInteraction["type"] {
			server.database.Query(context.Background(), `REMOVE @interactionKey IN Interactions`, map[string]interface{}{
				"interactionKey": metadata.Key,
			})
			json.NewEncoder(rw).Encode(map[string]string{
				"inserted": "n",
			})

		} else {
			server.database.Query(context.Background(), `UPDATE @interactionKey WITH { type: @type } IN Interactions`, map[string]interface{}{
				"interactionKey": readedInteraction["_key"],
				"type":           interaction,
			})
			json.NewEncoder(rw).Encode(map[string]string{
				"inserted": "c",
			})

		}

		return
	} else if err != nil {
		writeError(rw, "Error can not connect with database. "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(map[string]string{
		"inserted": "y",
	})

}

//postBright route: /brights
func (server *Server) postBright(rw http.ResponseWriter, r *http.Request) {

	println("Hola")
	_, username := middleware.AuthenticatedUser(r)

	err := r.ParseMultipartForm(24 >> 20)
	if err != nil {
		writeError(rw, "Problem parsing form", http.StatusInternalServerError)
		return
	}

	media := make([]string, 0, 4)
	for i, v := range r.MultipartForm.File {
		println(i)
		for _, h := range v {
			srcFile, err := h.Open()
			fmt.Println(h.Header)
			fmt.Println(err)
			dirPath := "/media/" + username + "/"
			os.MkdirAll(dirPath, 0755)
			dstFilepath := dirPath + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + i
			if strings.HasPrefix(h.Header.Get("Content-Type"), "video") {
				dstFilepath = dstFilepath + ".video"
			}
			dstFile, err := os.OpenFile(dstFilepath, os.O_CREATE|os.O_WRONLY, 0755)
			fmt.Println(err)
			io.Copy(dstFile, srcFile)

			srcFile.Close()
			dstFile.Close()

			media = append(media, dstFilepath)
		}
	}

	content := r.FormValue("content")

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
	_, username := middleware.AuthenticatedUser(r)

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
	for i, v := range r.MultipartForm.File {
		println(i)
		for _, h := range v {
			srcFile, err := h.Open()
			fmt.Println(err)
			dirPath := "/media/" + username + "/"
			os.MkdirAll(dirPath, 0755)
			dstFilepath := dirPath + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + i
			dstFile, err := os.OpenFile(dstFilepath, os.O_CREATE|os.O_WRONLY, 0755)
			fmt.Println(err)
			io.Copy(dstFile, srcFile)

			srcFile.Close()
			dstFile.Close()

			media = append(media, dstFilepath)
		}
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

	_, username := middleware.AuthenticatedUser(r)

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	// println(offset, limit)

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
		// fmt.Printf("%#v\n", u)
		if err != nil {
			continue
		}

		brights = append(brights, u)
	}

	json.NewEncoder(rw).Encode(brights)

}

package router

import (
	"brilla/internal/database/queries"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/arangodb/go-driver"
)

func (server *Server) getSearch(rw http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("q")
	aqlQuery := "FOR b IN Brillo\nLET f = @query\nFILTER "
	searchByTag := true

	for i, v := range strings.Fields(query) {
		if !strings.HasPrefix(v, "#") {
			aqlQuery = queries.SearchContentQuery
			searchByTag = false
			break
		}

		if i != 0 {
			aqlQuery = aqlQuery + " && "
		}
		aqlQuery = aqlQuery + "CONTAINS(b.content, \"" + v + "\")"
	}

	if searchByTag {
		aqlQuery = aqlQuery + " RETURN b"
	}

	if strings.HasPrefix(query, "@") {
		aqlQuery = queries.SearchUserQuery
		query = strings.TrimPrefix(query, "@")
	}

	println(aqlQuery)

	if query == "" {
		json.NewEncoder(rw).Encode(make([]bool, 0, 0))
		return
	}

	cursor, err := server.database.Query(driver.WithQueryCount(context.Background()), aqlQuery, map[string]interface{}{
		"query": query,
	})
	if err != nil {
		writeError(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close()

	results := make([]map[string]interface{}, 0, cursor.Count())

	for cursor.HasMore() {
		var result map[string]interface{}
		_, err := cursor.ReadDocument(context.Background(), &result)
		if err != nil {
			writeError(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, result)
	}

	json.NewEncoder(rw).Encode(results)

}

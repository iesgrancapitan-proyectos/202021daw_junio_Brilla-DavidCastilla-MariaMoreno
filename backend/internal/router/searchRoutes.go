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
	aqlQuery := queries.SearchContentQuery

	if strings.HasPrefix(query, "@") {
		aqlQuery = queries.SearchUserQuery
		query = strings.TrimPrefix(query, "@")
	}

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

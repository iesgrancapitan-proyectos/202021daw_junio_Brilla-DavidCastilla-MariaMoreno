package router

import (
	"net/http"

	"github.com/arangodb/go-driver"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router   *httprouter.Router
	database *driver.Database
}

var _ http.Handler = &Router{}

func New(database *driver.Database) *Router {

	router := httprouter.New()

	return &Router{
		router:   router,
		database: database,
	}

}

func routes(router *httprouter.Router) {
	// TODO: add routes
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	router.router.ServeHTTP(rw, r)
}

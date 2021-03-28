package router

import (
	"net/http"

	"github.com/arangodb/go-driver"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router   *httprouter.Router
	database *driver.Client
}

var _ http.Handler = &Router{}

func New(database *driver.Client) *Router {

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

package router

import (
	"net/http"

	"github.com/arangodb/go-driver"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router   *httprouter.Router
	database driver.Database
}

var _ http.Handler = &Server{}

func New(database driver.Database) *Server {

	router := httprouter.New()

	service := &Server{
		router:   router,
		database: database,
	}

	routes(service)

	return service

}

func routes(server *Server) {
	server.router.HandlerFunc(http.MethodPost, "/user", server.postUser)
	server.router.HandlerFunc(http.MethodGet, "/user/:username", server.getUser)
	server.router.HandlerFunc(http.MethodGet, "/user/:username/brights", server.getUserBrights)
	server.router.HandlerFunc(http.MethodPut, "/user/:username/follow", server.postUserFollow)
	server.router.HandlerFunc(http.MethodDelete, "/user/delete", server.deleteUser)
	server.router.HandlerFunc(http.MethodPost, "/user/login", server.postLogin)

	server.router.HandlerFunc(http.MethodPost, "/brights", server.postBright)
	server.router.HandlerFunc(http.MethodGet, "/brights/:id", server.getBright)
	server.router.HandlerFunc(http.MethodPost, "/brights/rebrilla", server.postRebrilla)
	server.router.HandlerFunc(http.MethodPost, "/brights/interaction", server.postInteraction)
	server.router.HandlerFunc(http.MethodPost, "/brights/comment", server.postComment)
	server.router.HandlerFunc(http.MethodDelete, "/brights/delete", server.deleteBright)
}

func (server *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(rw, r)
}

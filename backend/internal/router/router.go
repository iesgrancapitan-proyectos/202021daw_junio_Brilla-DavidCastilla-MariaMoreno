package router

import (
	"brilla/internal/middleware"
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
	server.router.HandlerFunc(http.MethodGet, "/brights/:id", server.getBright)
	server.router.HandlerFunc(http.MethodPost, "/user/login", server.postLogin)
	server.router.HandlerFunc(http.MethodPost, "/user/activate", server.postActivateUser)
	server.router.HandlerFunc(http.MethodGet, "/nfollowers/:username", server.getFollowers)
	server.router.HandlerFunc(http.MethodGet, "/nfollowed/:username", server.getFollowed)
	server.router.HandlerFunc(http.MethodGet, "/user/:username/brights/count", server.getNumBrillos)
	server.router.HandlerFunc(http.MethodPost, "/user/exits", server.getEmailExits)
	server.router.HandlerFunc(http.MethodPost, "/bright/ncomments", server.getNumComments)
	server.router.HandlerFunc(http.MethodPost, "/bright/nrebrillos", server.getNumRebrillos)
	server.router.HandlerFunc(http.MethodGet, "/auth/google", server.googleAuth)
	server.router.HandlerFunc(http.MethodGet, "/auth/google/confirm", server.googleAuthConfirm)

	server.router.HandlerFunc(http.MethodPut, "/user/:username/follow", middleware.NeedsAuth(server.database, server.putUserFollow))
	server.router.HandlerFunc(http.MethodDelete, "/user/delete", middleware.NeedsAuth(server.database, server.deleteUser))
	server.router.HandlerFunc(http.MethodPost, "/brights", middleware.NeedsAuth(server.database, server.postBright))
	server.router.HandlerFunc(http.MethodPost, "/brights/rebrilla", middleware.NeedsAuth(server.database, server.postRebrilla))
	server.router.HandlerFunc(http.MethodPost, "/brights/interaction", middleware.NeedsAuth(server.database, server.postInteraction))
	server.router.HandlerFunc(http.MethodPost, "/brights/comment", middleware.NeedsAuth(server.database, server.postComment))
	server.router.HandlerFunc(http.MethodDelete, "/brights/:idbrillo/delete", middleware.NeedsAuth(server.database, server.deleteBright))
	server.router.HandlerFunc(http.MethodGet, "/timeline", middleware.NeedsAuth(server.database, server.getTimeline))
	server.router.HandlerFunc(http.MethodGet, "/refresh", middleware.NeedsAuth(server.database, server.getRefresh))
	server.router.HandlerFunc(http.MethodGet, "/logout", middleware.NeedsAuth(server.database, server.getLogout))
	server.router.HandlerFunc(http.MethodGet, "/user/:username/isfollowing", middleware.NeedsAuth(server.database, server.isFollowing))

}

func (server *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")

	server.router.ServeHTTP(rw, r)
}

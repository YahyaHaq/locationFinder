package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/user", app.createUserHandler)
	router.HandlerFunc(http.MethodGet, "/user/:id", app.getUserHandler)
	router.HandlerFunc(http.MethodPatch, "/user/:id/location", app.updateLocationHandler)
	router.HandlerFunc(http.MethodGet, "/users/path", app.getUsersShortestPathHandler)
	router.HandlerFunc(http.MethodGet, "/users/nearby", app.getNearbyUsersHandler)

	return router
}

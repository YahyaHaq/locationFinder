package main

import (
	"fmt"
	"net/http"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create user")
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get user")
}

func (app *application) updateLocationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update location")
}

func (app *application) getUsersShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get shortest path")
}

func (app *application) getNearbyUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get nearby user")
}

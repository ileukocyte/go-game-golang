package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ileukocyte/go-game-golang/db"
)

func GameHandler(env *db.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Session accessed (id=%s)\n", mux.Vars(r)["id"])

		http.ServeFile(w, r, "assets/session.html")
	}
}

func NewGameHandler(env *db.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Initiating a new session...")
	}
}

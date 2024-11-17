package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gorilla/mux"
	"github.com/ileukocyte/go-game-golang/api"
	"github.com/ileukocyte/go-game-golang/db"
	"github.com/markbates/goth/gothic"
)

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		/*t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, gothUser)*/

		_, _ = fmt.Fprintf(w, "Hello, %s!", gothUser.Name)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)

	if err != nil {
		_, _ = fmt.Fprint(w, err)

		return
	}

	_, _ = fmt.Fprintf(w, "Logged in as: %s", user.Name)
}

func main() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	env, err := db.InitDatabase(connStr)

	if err != nil {
		log.Fatal(err)
	}

	googleConfig := &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_CALLBACK_URL"),
		Scopes:       []string{"email", "profile", "openid"},
		Endpoint:     google.Endpoint,
	}

	router := mux.NewRouter().StrictSlash(true)

	var (
		cssFs = http.FileServer(http.Dir("assets/css/"))
		jsFs  = http.FileServer(http.Dir("assets/js/"))
	)

	router.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", cssFs))
	router.PathPrefix("/assets/js/").Handler(http.StripPrefix("/assets/js/", jsFs))

	router.HandleFunc("/game/new", api.NewGameHandler(env))
	router.HandleFunc("/game/{id:[0-9]+}", api.GameHandler(env))

	router.HandleFunc("/auth", api.AuthHandler)
	router.HandleFunc("/auth/oauth", api.OAuthHandler(googleConfig))
	router.HandleFunc("/auth/google/callback", api.GoogleCallbackHandler(env, googleConfig))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router))
}

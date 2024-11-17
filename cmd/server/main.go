package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

func newGameHandler(e *db.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Initiating a new session...")
	}
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

	/*goth.UseProviders(
		google.New(
			os.Getenv("OAUTH_CLIENT_ID"),
			os.Getenv("OAUTH_CLIENT_SECRET"),
			os.Getenv("OAUTH_CALLBACK_URL"),
		),
	)*/

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/game/new", newGameHandler(env))

	var (
		cssFs = http.FileServer(http.Dir("assets/css/"))
		jsFs  = http.FileServer(http.Dir("assets/js/"))
	)

	router.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", cssFs))
	router.PathPrefix("/assets/js/").Handler(http.StripPrefix("/assets/js/", jsFs))

	router.HandleFunc("/game/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("Session ID: %s", mux.Vars(r)["id"])

		http.ServeFile(w, r, "assets/session.html")
	})

	//router.HandleFunc("/auth/{provider}", signInHandler)
	//router.HandleFunc("/auth/{provider}/callback", callbackHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router))
}

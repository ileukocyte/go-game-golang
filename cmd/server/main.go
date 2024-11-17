package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
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
	goth.UseProviders(
		google.New(
			os.Getenv("OAUTH_CLIENT_ID"),
			os.Getenv("OAUTH_CLIENT_SECRET"),
			os.Getenv("OAUTH_CALLBACK_URL"),
		),
	)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/session/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/session.html")
	})
	//router.HandleFunc("/auth/{provider}", signInHandler)
	//router.HandleFunc("/auth/{provider}/callback", callbackHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), router))
}

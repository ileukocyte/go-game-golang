package api

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/ileukocyte/go-game-golang/db"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/auth.html")
}

func OAuthHandler(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func GoogleCallbackHandler(env *db.Env, config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		t, err := config.Exchange(context.Background(), code)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		// TODO: handle signing up and logging in
		client := config.Client(context.Background(), t)
		res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		_ = res
	}
}

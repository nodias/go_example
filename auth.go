package main

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

const (
	nextPageKey     = "next_page"
	authSecurityKey = "auth_security_key"
)

func init() {
	gomniauth.SetSecurityKey(authSecurityKey)
	gomniauth.WithProviders(
		google.New("636296155193-a9abes4mc1p81752l116qkr9do6oev3f.apps.googleusercontent.com", "EVvuy0Agv4jWflml0pvC6-vI", "http://127.0.0.1:3000/auth/callback/google"),
	)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	action := ps.ByName("action")
	provider := ps.ByName("provider")
	s := sessions.GetSession(r)
}

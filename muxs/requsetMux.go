package muxs

import (
	"authenticate/handler"
	"net/http"
)

func GetRequesttMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", handler.LoginHandler)
	mux.HandleFunc("/auth/register", handler.RegisterHandler)
	mux.HandleFunc("/auth/logout", handler.LogoutHandler)
	return mux
}

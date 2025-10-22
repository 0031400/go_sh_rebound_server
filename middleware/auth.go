package middleware

import (
	"go_sh_rebound_server/config"
	"net/http"
)

func AuthNodeMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("authorization")
		if auth != config.NodeAuth {
			return
		}
		next(w, r)
	}
}

func AuthClientMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("authorization")
		if auth != config.ClientAuth {
			return
		}
		next(w, r)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// Eine Testmiddleware Vorlage
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection von allen POST Requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// Speichert und ladet alle Sessions von den Seiten
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
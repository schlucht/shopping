package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/booking/pkg/config"
	"github.com/schlucht/booking/pkg/handlers"
	"github.com/schlucht/booking/pkg/render"
)

const PORTNUMBER = ":8080"
var app config.AppConfig
// var session *scs.SessionManager

var session *scs.SessionManager	
func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session	= session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create Template")
	}

	app.TemplateCatche = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Println("Server run on localhost:8080")

	srv := &http.Server{
		Addr: PORTNUMBER,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

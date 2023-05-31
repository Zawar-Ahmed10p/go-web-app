package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Zawar-Ahmed10p/go-web-app/pkg/config"
	"github.com/Zawar-Ahmed10p/go-web-app/pkg/handlers"
	"github.com/Zawar-Ahmed10p/go-web-app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var appConf config.Appconfig
var session *scs.SessionManager

func main() {

	appConf.InProduction = false // for prod swithc to true
	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConf.InProduction
	appConf.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache.", err)
	}
	appConf.TemplateCache = tc
	appConf.UseCache = false
	repo := handlers.NewRepo((&appConf))
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConf)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	log.Printf("Starting server on %+s", portNumber)
	// http.ListenAndServe(portNumber, nil)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConf),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}

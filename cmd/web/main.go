package main

import (
	"log"
	"net/http"

	"github.com/Zawar-Ahmed10p/go-web-app/pkg/config"
	"github.com/Zawar-Ahmed10p/go-web-app/pkg/handlers"
	"github.com/Zawar-Ahmed10p/go-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var appConf config.Appconfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache.")
	}
	appConf.TemplateCache = tc
	appConf.UseCache = false
	repo := handlers.NewRepo((&appConf))
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConf)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	log.Printf("Starting server on %+s", portNumber)
	http.ListenAndServe(portNumber, nil)
}

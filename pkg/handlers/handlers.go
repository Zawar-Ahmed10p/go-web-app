package handlers

import (
	"net/http"

	"github.com/Zawar-Ahmed10p/go-web-app/pkg/config"
	"github.com/Zawar-Ahmed10p/go-web-app/pkg/models"
	"github.com/Zawar-Ahmed10p/go-web-app/pkg/render"
)

type Repository struct {
	App *config.Appconfig
}

var Repo *Repository

//create new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

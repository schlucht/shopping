package handlers

import (
	"net/http"

	"github.com/schlucht/booking/pkg/config"
	"github.com/schlucht/booking/pkg/models"
	"github.com/schlucht/booking/pkg/render"
)

// TemplateData holds data send from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hallo Lothar"

	remotIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remotIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["name"] = "Schmid Lothar"
	stringMap["mail"] = "jagolo@jagolo.ch"

	remotIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remotIP
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Avalability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-avalaible.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

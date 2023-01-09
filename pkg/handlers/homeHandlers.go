package handlers

import (
	"net/http"

	"github.com/schlucht/booking/pkg/models"
	"github.com/schlucht/booking/pkg/render"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{}, r)

}

package handlers

import (
	"net/http"

	"github.com/schlucht/booking/pkg/models"
	"github.com/schlucht/booking/pkg/render"
)

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hallo Lothar"

	remotIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remotIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	}, r)
}

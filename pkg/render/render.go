package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/schlucht/booking/pkg/config"
	"github.com/schlucht/booking/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig
var folders = []string{"components", "layouts"}

func NewTemplate(a *config.AppConfig) {
	app = a
}

func DefaultData() {

}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate render Website with Data from Template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData, r *http.Request) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCatche
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache ", tmpl)
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return
	}
}

// CreateTemplateCache creates  a template cache a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles((page))
		if err != nil {
			return myCache, err
		}
		for _, folder := range folders {
			matches, err := filepath.Glob("./templates/" + folder + "/*.tmpl")
			if err != nil {
				return myCache, err
			}
			if len(matches) > 0 {
				ts, err = ts.ParseGlob("./templates/" + folder + "/*.tmpl")
				if err != nil {
					log.Println(err)
					return myCache, err
				}
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

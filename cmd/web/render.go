package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template

	// if we are using the template cache. try to get the template from our map, stored in the recent.
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateLateFromDisk(t)
		if err != nil {
			log.Println("Error building template from disk:", err)
			return
		}

		log.Println("build template from disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) buildTemplateLateFromDisk(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/footer.partial.gohtml",
		"./templates/partials/header.partial.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	if app.templateMap == nil {
		app.templateMap = make(map[string]*template.Template)
	}

	app.templateMap[t] = tmpl

	return tmpl, nil
}

package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/lucasleaopinto/bookings/internal/config"
	"github.com/lucasleaopinto/bookings/internal/models"
)

var app *config.AppConfig

// NewTemplates sets the confog for the template packaage
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate render templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	//create a template cache
	//fmt.Println("entrou no render 01 ")

	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			fmt.Println("error create template into UseCache false.", err)
			log.Println(err)
		}

	}

	//tc, err := CreateTemplateCache()
	// if err != nil {
	// 	fmt.Println("error RenderTemplate: 01 ", err)
	// 	log.Fatal(err)
	// }

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache.")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err = t.Execute(buf, td)
	if err != nil {
		fmt.Println("error executing buffer. ", err)
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writting buffer. ", err)
		log.Println(err)
		//fmt.Println("error parsing template: ", err)
		return
	}
}

// CreateTemplateCache create templates using html/template
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmp from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			fmt.Printf("page error:%s ", page)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	if len(myCache) == 0 {
		fmt.Println("no template loaded ")
		//TO DO
		//criar um erro customizado
	}

	return myCache, nil

}

// var tc = make(map[string]*template.Template)

// //RenderTemplate render templates using html/template and add to cache
// func RenderTemplate(w http.ResponseWriter, t string) {

// 	var err error
// 	var tmpl *template.Template

// 	//check is exist a template in cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cache")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil

// }

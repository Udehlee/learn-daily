package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

// init initialises the templates
func init() {

	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// index renders the "index.gohtml" template.
func index(res http.ResponseWriter, req *http.Request) {

	err := tmpl.ExecuteTemplate(res, "index.gohtml", nil)

	if err != nil {
		log.Fatalln("template didnt execute", nil)
	}

}

func main() {

	http.HandleFunc("/", index)

	// Serve static files from the "public" directory.

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)

}

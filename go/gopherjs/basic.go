package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Encounter error while parsing template:", err)
	}

	t.Execute(w, templateData)
}

func BasicHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "./templates/basic.html", nil)
	})
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", BasicHandler()).Methods("GET")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}

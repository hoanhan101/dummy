package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Gopher struct {
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Gopher"
	}
	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./templates/sample.html", gopher)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello-gopher", helloHandler)
	http.Handle("/", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":8080", nil)
}

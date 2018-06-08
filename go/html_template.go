package main

import (
	"html/template"
	"log"
	"net/http"
)

type Gopher struct {
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    log.Print("/hello-gopher")

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
    log.Print("Server is running at port 8080")

    http.HandleFunc("/hello-gopher", helloHandler)
    http.ListenAndServe(":8080", nil)
}

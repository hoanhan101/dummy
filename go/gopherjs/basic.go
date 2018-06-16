package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

func LowercaseHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s string

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Encountered error while reading request body:", err)
		}

		bodyString := string(body)
		log.Println("Get request from client:", bodyString)

		// Unmarshal the request string.
		err = json.Unmarshal([]byte(bodyString), &s)
		if err != nil {
			log.Println("Encoutner error while unmarshalling json", err)
		}

		// Convert to lowercase, marshall it, then send it back.
		textBytes, err := json.Marshal(strings.ToLower(s))
		if err != nil {
			log.Println("Encoutner error while marshalling json", err)
		}

		log.Println("Finish lowercase task then send back:", string(textBytes))
		w.Write(textBytes)
	})
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", BasicHandler()).Methods("GET")
	r.Handle("/lowercase", LowercaseHandler())

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}

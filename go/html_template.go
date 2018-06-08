package main

import (
	"html/template"
	"log"
	"net/http"
)

type Gopher struct {
	Name string
}

func hello(w http.ResponseWriter, r *http.Request) {
	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Gopher"
	}
	gopher := Gopher{Name: gophername}
	re
}

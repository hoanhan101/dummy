package main

import (
	"log"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func DisplayAlertMessageJSGlobal(message string) {
	log.Println("DisplayAlertMessageJSGlobal")
	js.Global.Call("alert", message)
}

func DisplayAlertMessageDOM(message string) {
	log.Println("DisplayAlertMessageDOM")
	dom.GetWindow().Alert(message)
}

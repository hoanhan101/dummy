package main

import (
	"fmt"

	"honnef.co/go/js/dom"
)

var d = dom.GetWindow().Document().(dom.HTMLDocument)

func run() {
	messageInput := d.GetElementByID("messageInput").(*dom.HTMLInputElement)

	alertButtonJS := d.GetElementByID("alertMessageJSGlobal").(*dom.HTMLButtonElement)
	alertButtonJS.AddEventListener("click", false, func(event dom.Event) {
		DisplayAlertMessageJSGlobal(messageInput.Value)
	})

	alertButtonDOM := d.GetElementByID("alertMessageDOM").(*dom.HTMLButtonElement)
	alertButtonDOM.AddEventListener("click", false, func(event dom.Event) {
		DisplayAlertMessageDOM(messageInput.Value)
	})

	lowercaseText := d.GetElementByID("lowercaseText").(*dom.HTMLButtonElement)
	lowercaseText.AddEventListener("click", false, func(event dom.Event) {
		go lowercaseTextTransformer()
	})
}

func main() {
	switch readyState := d.ReadyState(); readyState {
	case "loading":
		d.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		panic(fmt.Errorf("Internal error: unexpected document.ReadyState value: %v", readyState))
	}
}

package main

import (
	"encoding/json"
	"log"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

func lowercaseTextTransformer() {
	// Get the text input text.
	d := dom.GetWindow().Document()
	textToLowerCase := d.GetElementByID("textToLowercase").(*dom.HTMLInputElement)

	log.Println("Get user input, prepare to send to server:", textToLowerCase.Value)

	// Marshall it, then send it to the POST endpoint.
	textBytes, err := json.Marshal(textToLowerCase.Value)
	if err != nil {
		log.Println("Encountered error while marshalling json", err)
	}

	data, err := xhr.Send("POST", "/lowercase", textBytes)
	if err != nil {
		log.Println("Encountered error while sending XHR POST:", err)
	}

	// After server send back the data, unmarshall it and display it using js
	// packages.
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Println("Encounter error while unmarshalling json:", err)
	}

	log.Println("Get back from server, unmarshal then display:", s)
	textToLowerCase.Set("value", s)
}

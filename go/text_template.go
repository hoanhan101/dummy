package main

import (
	// "text/template"
	"fmt"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{Material: "wood", Count: 17}
	fmt.Println(sweaters)
}

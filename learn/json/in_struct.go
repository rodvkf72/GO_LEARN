package main

import (
	"encoding/json"
	"fmt"
)

type Fields struct {
	VisibleField   string `json:"visibleField"`
	InvisibleField string `json:"invisibleField"`
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:"invisibleField,omitempty"` // <-- exclude Field
		Additional     string `json:"additional"`
		Additional2    string `json:"additional2"`
	}{Fields: f, Additional: "c", Additional2: "d"})
	fmt.Println(string(b))
}

func main() {
	ExampleOmitFields()
}

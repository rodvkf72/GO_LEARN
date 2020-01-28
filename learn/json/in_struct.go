package main

import (
	"encoding/json"
	"fmt"
)

type Fields struct {
	VisibleField   string
	InvisibleField string
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:",omitempty"` // <-- exclude Field
		Additional     string
		Additional2    string
	}{Fields: f, Additional: "c", Additional2: "d"})
	fmt.Println(string(b))
}

func main() {
	ExampleOmitFields()
}

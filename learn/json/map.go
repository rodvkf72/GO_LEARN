package main

import (
	"encoding/json"
	"fmt"
)

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{
		"Name": "Kim",
		"Age":  26,
	})
	fmt.Println(string(b))
}

func main() {
	Example_mapMarshalJSON()
}

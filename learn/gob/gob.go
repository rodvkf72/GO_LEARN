package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("% x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
}

func main() {
	Example_gob()
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello World !")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

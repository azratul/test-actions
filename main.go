package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello desde Go en K3d sobre RPi4!")
	})

	log.Println("Server corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

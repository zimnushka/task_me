package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/apps", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "<h1>Hello world</h1>")

		fmt.Println("their print")
	})

	http.ListenAndServe(":8080", nil)
}

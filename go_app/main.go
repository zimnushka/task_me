package main

import (
	"bytes"
	"fmt"
	"net/http"

	appControllers "github.com/zimnushka/task_me_go/go_app/controllers"
)

func main() {

	http.HandleFunc("/apps", func(w http.ResponseWriter, _ *http.Request) {
		apps := appControllers.GetApps()
		var buffer bytes.Buffer

		for _, app := range apps {
			buffer.WriteString(app.Name)
		}

		fmt.Fprintf(w, "<h1>"+buffer.String()+"</h1>")

	})

	http.ListenAndServe(":8080", nil)
}

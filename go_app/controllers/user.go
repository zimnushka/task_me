package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

func UserController(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	users := repositories.GetUsers()
	var buffer bytes.Buffer

	for _, user := range users {
		buffer.WriteString(user.Name)
	}

	fmt.Fprintf(w, "<h1>"+buffer.String()+"</h1>")
}

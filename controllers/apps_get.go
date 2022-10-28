package apps

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getApps() []App {
	db, err := sql.Open("mysql", "root:43WYOH5l8W1I@tcp(192.168.17.9:3306)/taskMe")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM users")

	if err != nil {
		panic(err.Error())
	}
	appsLng := 0
	apps := make([]App, appsLng)

	for results.Next() {
		var app App
		err := results.Scan(&app.Id, &app.Name, &app.Password, &app.Email)
		if err != nil {
			panic(err.Error())
		}
		apps = append(apps, app)
		appsLng++
	}

	defer results.Close()
	return apps
}

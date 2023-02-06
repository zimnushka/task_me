package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TaskMeDB struct {
}

func (TaskMeDB) GetDB() (*sql.DB, error) {
	const user = "root"
	const password = "43WYOH5l8W1I"
	const url = "mariadb:3306"
	const db = "taskMe"

	const debugUrl = "localhost:3306"

	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, debugUrl, db))
}

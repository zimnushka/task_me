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
	const url = "192.168.17.9:3306"
	const db = "taskMe"

	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, db))
}

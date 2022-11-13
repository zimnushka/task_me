package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type TaskMeDB struct {
}

func (TaskMeDB) GetDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:43WYOH5l8W1I@tcp(192.168.17.9:3306)/taskMe")
}

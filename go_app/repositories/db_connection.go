package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/app"
)

type TaskMeDB struct {
}

func (TaskMeDB) GetDB() (*sql.DB, error) {
	config := app.GetConfig()

	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DBParams.User, config.DBParams.Password, config.DBParams.Url, config.DBParams.Db))
}

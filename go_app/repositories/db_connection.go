package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() sql.DB {
	db, err := sql.Open("mysql", "root:43WYOH5l8W1I@tcp(192.168.17.9:3306)/taskMe")
	if err != nil {
		panic(err.Error())
	}
	return *db
}
func queryDB(db sql.DB, query string) *sql.Rows {
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	return results
}

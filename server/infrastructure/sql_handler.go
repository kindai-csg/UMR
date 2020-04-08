package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type SqlHandler struct {
	db *sql.DB
}

func NewSqlHandler() *SqlHandler {
	db, err := sql.Open("mysql", "root:densan@tcp(mysql)/umr")
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	sqlHandler := SqlHandler {
		db: db,
	}
	return &sqlHandler
}

func (handler *SqlHandler) Query(query string) (*sql.Rows, error){
	rows, err := handler.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

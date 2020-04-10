package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type SqlHandler struct {
	db *sql.DB
}

type SqlConfig struct {
	User string  `toml:"user"`
	Password string  `toml:"password"`
	Host string  `toml:"host"`
	Port string  `toml:"port"`
	Database string  `toml:"database"`
}

func NewSqlHandler(config SqlConfig) *SqlHandler {
	db, err := sql.Open("mysql", config.User+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database)
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

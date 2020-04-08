package database

import (
	"database/sql"
)

type SqlHandler interface {
	Query(query string) (*sql.Rows, error)
}
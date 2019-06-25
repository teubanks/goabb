package store

import "database/sql"

type Store struct {
	dbConn *sql.DB
}

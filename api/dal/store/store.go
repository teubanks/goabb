package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// Store holds everything needed to communicate with a database. It implements
// the Dal interface
type Store struct {
	dbConn *sql.DB
}

type Filters map[string]interface{}

func NewStore(user, pass, pgUrl string) *Store {
	connStr := fmt.Sprintf("user=%s dbname=%s Host=%s sslmode=disable", user, pass, pgUrl)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	s := &Store{db}
	return s
}

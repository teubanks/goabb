package store

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Store holds everything needed to communicate with a database. It implements
// the Dal interface
type Store struct {
	dbConn *sqlx.DB
}

type Filters map[string]interface{}

// NewStore returns an pointer to an instance of Store with a DB connection
// established
func NewStore(user, pass, dbName, pgUrl string) *Store {
	connStr := fmt.Sprintf("user=%s password=%s database=%s host=%s sslmode=disable", user, pass, dbName, pgUrl)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	s := &Store{db}
	return s
}

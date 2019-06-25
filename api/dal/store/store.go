package store

import "database/sql"

// Store holds everything needed to communicate with a database. It implements
// the Dal interface
type Store struct {
	dbConn *sql.DB
}

type Filters map[string]interface{}

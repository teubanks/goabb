package store

import (
	"github.com/teubanks/goabb/api/dal/models"
)

var (
	getUsersStmt   = "SELECT * FROM users"
	insertUserStmt = "INSERT INTO users (first_name, last_name, email, age, birthday) VALUSE (:first_name, :last_name, :email, :age, :birthday)"
	deleteUserStmt = "DELETE FROM users WHERE id=?"
)

// ReadUsers queries the database for all users
func (s *Store) ReadUsers() ([]*models.User, error) {
	var users []*models.User
	err := s.dbConn.Select(users, getUsersStmt)
	return users, err
}

// ReadUser queries the database for the user that exists for the passed in ID
func (s *Store) ReadUser(id int) (*models.User, error) {
	filters := make(Filters)
	filters["id"] = id
	return s.readUser(filters)
}

// ReadUserByEmail queries the database for the user that exists with the
// specified email
func (s *Store) ReadUserByEmail(email string) (*models.User, error) {
	filters := make(Filters)
	filters["email"] = email
	return s.readUser(filters)
}

func (s *Store) readUser(filters Filters) (*models.User, error) {
	var user *models.User

	clauses := []string{}
	// for k, v := range filters {
	// 	clauses := fmt.Sprintf("%s=?", k, )
	// }
	if len(clauses) == 0 {
		return user, nil
	}

	// query := fmt.Sprintf(getUsersStmt + " WHERE "
	return &models.User{}, nil
}

// CreateUser inserts a new user into the system
func (s *Store) CreateUser(u *models.User) error {
	_, err := s.dbConn.NamedExec(insertUserStmt, u)
	return err
}

// UpdateUser updates the passed in models.User using the members of the passed
// in struct as attributes
func (s *Store) UpdateUser(u *models.User) error {
	return nil
}

// DeleteUser deletes the user that has the specified ID
func (s *Store) DeleteUser(id int) error {
	_, err := s.dbConn.Exec(deleteUserStmt, id)
	return err
}

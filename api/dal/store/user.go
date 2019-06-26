package store

import "github.com/outdoorsy/api/model"

var (
	getUsersStmt = "SELECT * FROM users"
 insertUserStmt = "INSERT INTO users (first_name, last_name, email, age, birthday) VALUSE (:first_name, :last_name, :email, :age, :birthday)"
deleteUserStmt = "DELETE FROM users WHERE id=?"
)
// ReadUsers queries the database for all users
func (s *Store) ReadUsers() ([]*model.User, error) {
	var users []*model.User{}
	err := s.dbConn.Select(users, getUsersStmt)
	return users, err
}

// ReadUser queries the database for the user that exists for the passed in ID
func (s *Store) ReadUser(id int) (*model.User, error) {
	filters := make(Filters)
	filters["id"] = id
	return s.readUser(filters)
}

// ReadUserByEmail queries the database for the user that exists with the
// specified email
func (s *Store) ReadUserByEmail(email string) (*model.User, error) {
	filters := make(Filters)
	filters["email"] = email
	return s.readUser(filters)
}

func (s *Store) readUser(filters map[string]interface{}) (*model.User, error) {
	var user *model.User

	clauses := []string{}
	for k, v := range filters {
		clauses := fmt.Sprintf("k=?")
	}
	if len(clauses) == 0 {
		return user, nil
	}

	query := fmt.Sprintf(getUsersStmt + " WHERE "
	return &model.User{}, nil
}

// CreateUser inserts a new user into the system
func (s *Store) CreateUser(u *model.User) error {
	return s.dbConn.NamedExec(insertUserStmt, u)
}

// UpdateUser updates the passed in model.User using the members of the passed
// in struct as attributes
func (s *Store) UpdateUser(u *model.User) error {
	return nil
}

// DeleteUser deletes the user that has the specified ID
func (s *Store) DeleteUser(u *model.User) error {
	return s.dbConn.Exec(deleteUserStmt, u.ID)
}

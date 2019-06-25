package store

import "github.com/outdoorsy/api/model"

// ReadUsers queries the database for all users
func (s *Store) ReadUsers() ([]*model.User, error) {
	return []*model.User{}, nil
}

// ReadUser queries the database for the user that exists for the passed in ID
func (s *Store) ReadUser(id int) (*model.User, error) {
	return &model.User{}, nil
}

// ReadUserByEmail queries the database for the user that exists with the
// specified email
func (s *Store) ReadUserByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}

// CreateUser inserts a new user into the system
func (s *Store) CreateUser(u *model.User) error {
	return nil
}

// UpdateUser updates the passed in model.User using the members of the passed
// in struct as attributes
func (s *Store) UpdateUser(u *model.User) error {
	return nil
}

// DeleteUser deletes the user that has the specified ID
func (s *Store) DeleteUser(id int) error {
	return nil
}

package store

import "github.com/teubanks/goabb/api/dal/models"

// MockStore should be used in tests only
type MockStore struct {
	OnReadUsers       func() ([]*models.User, error)
	OnReadUser        func(int) (*models.User, error)
	OnReadUserByEmail func(string) (*models.User, error)
	OnCreateUser      func(*models.User) error
	OnUpdateUser      func(*models.User) error
	OnDeleteUser      func(int) error
}

// ReadUsers will return the results of OnReadUsers if set. Otherwise, an empty
// slice of models.Users and a nil error will be returned
func (m *MockStore) ReadUsers() ([]*models.User, error) {
	if m.OnReadUsers != nil {
		return m.OnReadUsers()
	}

	return []*models.User{}, nil
}

// ReadUser will return the results of OnReadUser, if set. Otherwise, an empty
// models.User struct and a nil error will be returned
func (m *MockStore) ReadUser(id int) (*models.User, error) {
	if m.OnReadUser != nil {
		return m.OnReadUser(id)
	}

	return &models.User{}, nil
}

// ReadUserByEmail will return the results of OnReadUserByEmail, if set.
// Otherwise, an empty models.User struct and a nil error will be returned
func (m *MockStore) ReadUserByEmail(email string) (*models.User, error) {
	if m.OnReadUserByEmail != nil {
		return m.OnReadUserByEmail(email)
	}

	return &models.User{}, nil
}

// CreateUser will return the result of OnCreateUser if set, otherwise, a nil
// error is returned
func (m *MockStore) CreateUser(u *models.User) error {
	if m.OnCreateUser != nil {
		return m.OnCreateUser(u)
	}
	return nil
}

// UpdateUser will return the result of OnUpdateUser if set, otherwise, a nil
// error is returned
func (m *MockStore) UpdateUser(u *models.User) error {
	if m.OnUpdateUser != nil {
		return m.OnUpdateUser(u)
	}
	return nil
}

// DeleteUser will return the result of OnDeleteUser if set, otherwise, a nil
// error is returned
func (m *MockStore) DeleteUser(id int) error {
	if m.OnDeleteUser != nil {
		return m.OnDeleteUser(id)
	}
	return nil
}

package store

import "github.com/outdoorsy/api/model"

// MockStore should be used in tests only
type MockStore struct {
	OnReadUsers       func() ([]*model.User, error)
	OnReadUser        func(int) (*model.User, error)
	OnReadUserByEmail func(string) (*model.User, error)
	OnCreateUser      func(*model.User) error
	OnUpdateUser      func(*model.User) error
	OnDeleteUser      func(int) error
}

// ReadUsers will return the results of OnReadUsers if set. Otherwise, an empty
// slice of model.Users and a nil error will be returned
func (m *MockStore) ReadUsers() ([]*model.User, error) {
	if m.OnReadUsers != nil {
		return m.OnReadUsers()
	}

	return []*model.User{}, nil
}

// ReadUser will return the results of OnReadUser, if set. Otherwise, an empty
// model.User struct and a nil error will be returned
func (m *MockStore) ReadUser(id int) (*model.User, error) {
	if m.OnReadUser != nil {
		return m.OnReadUser(id)
	}

	return &model.User{}, nil
}

// ReadUserByEmail will return the results of OnReadUserByEmail, if set.
// Otherwise, an empty model.User struct and a nil error will be returned
func (m *MockStore) ReadUserByEmail(email string) (*model.User, error) {
	if m.OnReadUserByEmail != nil {
		return m.OnReadUserByEmail(email)
	}

	return &model.User{}, nil
}

// CreateUser will return the result of OnCreateUser if set, otherwise, a nil
// error is returned
func (m *MockStore) CreateUser(u *model.User) error {
	if m.OnCreateUser != nil {
		return m.OnCreateUser(u)
	}
	return nil
}

// UpdateUser will return the result of OnUpdateUser if set, otherwise, a nil
// error is returned
func (m *MockStore) UpdateUser(u *model.User) error {
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

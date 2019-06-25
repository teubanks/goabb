package dal

import "github.com/outdoorsy/api/model"

type Dal interface {
	GetUsers() ([]*model.User, error)
	GetUser(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(id int) error
}

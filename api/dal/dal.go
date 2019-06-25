package dal

import "github.com/outdoorsy/api/model"

type Dal interface {
	ReadUsers() ([]*model.User, error)
	ReadUser(id int) (*model.User, error)
	ReadUserByEmail(email string) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	DeleteUser(id int) error
}

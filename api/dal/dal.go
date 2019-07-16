package dal

import "github.com/teubanks/goabb/api/dal/models"

type Dal interface {
	ReadUsers() ([]*models.User, error)
	ReadUser(id int) (*models.User, error)
	ReadUserByEmail(email string) (*models.User, error)
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	DeleteUser(id int) error
}

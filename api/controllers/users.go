package goabb

import (
	"context"
	"log"

	"github.com/teubanks/goabb/api/dal"
	users "github.com/teubanks/goabb/api/gen/users"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type usersSrvc struct {
	logger *log.Logger
	db     dal.Dal
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger, db dal.Dal) users.Service {
	return &usersSrvc{logger, db}
}

// Create implements create.
func (s *usersSrvc) Create(ctx context.Context, p *users.UserCreatePayload) (res *users.User, err error) {
	res = &users.User{}
	s.logger.Print("users.create")
	return
}

// Read implements read.
func (s *usersSrvc) Read(ctx context.Context, p *users.ReadPayload) (res *users.User, err error) {
	res = &users.User{}
	s.logger.Print("users.read")
	return
}

func (s *usersSrvc) ReadAll(ctx context.Context) (res []*users.User, err error) {
	res = []*users.User{}
	s.logger.Print("users.read_all")
	return
}

func (s *usersSrvc) Delete(ctx context.Context, p *users.DeletePayload) error {
	return nil
}

func (s *usersSrvc) Update(ctx context.Context, p *users.UserUpdatePayload) error {
	return nil
}

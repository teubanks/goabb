package goabb

import (
	"context"
	"log"

	users "github.com/teubanks/goabb/api/gen/users"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &userssrvc{logger}
}

// Create implements create.
func (s *userssrvc) Create(ctx context.Context, p *users.UserCreatePayload) (res *users.User, err error) {
	res = &users.User{}
	s.logger.Print("users.create")
	return
}

// Read implements read.
func (s *userssrvc) Read(ctx context.Context, p *users.ReadPayload) (res *users.User, err error) {
	res = &users.User{}
	s.logger.Print("users.read")
	return
}

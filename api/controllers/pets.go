package goabb

import (
	"context"
	"log"

	pets "github.com/teubanks/goabb/gen/pets"
)

// pets service example implementation.
// The example methods log the requests and return zero values.
type petssrvc struct {
	logger *log.Logger
}

// NewPets returns the pets service implementation.
func NewPets(logger *log.Logger) pets.Service {
	return &petssrvc{logger}
}

// Create implements create.
func (s *petssrvc) Create(ctx context.Context, p *pets.PetCreatePayload) (res *pets.Pet, err error) {
	res = &pets.Pet{}
	s.logger.Print("pets.create")
	return
}

// Read implements read.
func (s *petssrvc) Read(ctx context.Context, p *pets.ReadPayload) (res *pets.Pet, err error) {
	res = &pets.Pet{}
	s.logger.Print("pets.read")
	return
}

// ReadAll implements read_all.
func (s *petssrvc) ReadAll(ctx context.Context) (res []*pets.Pet, err error) {
	s.logger.Print("pets.read_all")
	return
}

// Update implements update.
func (s *petssrvc) Update(ctx context.Context, p *pets.PetUpdatePayload) (err error) {
	s.logger.Print("pets.update")
	return
}

// Delete implements delete.
func (s *petssrvc) Delete(ctx context.Context, p *pets.DeletePayload) (err error) {
	s.logger.Print("pets.delete")
	return
}

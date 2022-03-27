package queries

import (
	"context"
	"petstoreproject/internal/petstore/domain"
)

type ListPetsHandler struct {
	petRepo domain.PetRepository
}

func NewListPetsHandler(petRepo domain.PetRepository) ListPetsHandler {
	if petRepo == nil {
		panic("nil petRepo")
	}

	return ListPetsHandler{petRepo: petRepo}
}

func (h ListPetsHandler) Handle(ctx context.Context) ([]domain.Pet, error) {
	return h.petRepo.ListPets(ctx)
}

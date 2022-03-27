package queries

import (
	"context"
	"petstoreproject/internal/petstore/domain"
)

type GetPetHandler struct {
	petRepo domain.PetRepository
}

func NewGetPetHandler(petRepo domain.PetRepository) GetPetHandler {
	if petRepo == nil {
		panic("nil petRepo")
	}

	return GetPetHandler{petRepo: petRepo}
}

func (h GetPetHandler) Handle(ctx context.Context, id string) (domain.Pet, error) {
	return h.petRepo.GetPet(ctx, id)
}

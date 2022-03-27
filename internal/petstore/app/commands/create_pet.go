package commands

import (
	"context"
	"petstoreproject/internal/petstore/domain"
)

type CreatePetHandler struct {
	petRepo domain.PetRepository
}

func NewCreatePetHandler(petRepo domain.PetRepository) CreatePetHandler {
	if petRepo == nil {
		panic("nil petRepo")
	}

	return CreatePetHandler{petRepo: petRepo}
}

func (h CreatePetHandler) Handle(ctx context.Context, pet domain.Pet) error {
	return h.petRepo.CreatePet(ctx, pet)
}

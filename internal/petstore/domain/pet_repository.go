package domain

import "context"

type PetRepository interface {
	ListPets(ctx context.Context) ([]Pet, error)
	CreatePet(ctx context.Context, pet Pet) error
	GetPet(ctx context.Context, id string) (Pet, error)
}

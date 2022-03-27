package app

import (
	"petstoreproject/internal/petstore/app/commands"
	"petstoreproject/internal/petstore/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatePet commands.CreatePetHandler
}

type Queries struct {
	ListPets queries.ListPetsHandler
	GetPet   queries.GetPetHandler
}

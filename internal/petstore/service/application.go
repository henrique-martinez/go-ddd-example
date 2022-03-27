package service

import (
	"context"
	"database/sql"
	"petstoreproject/internal/petstore/adapters"
	"petstoreproject/internal/petstore/app"
	"petstoreproject/internal/petstore/app/commands"
	"petstoreproject/internal/petstore/app/queries"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func NewApplication(ctx context.Context) app.Application {
	connString := "postgresql://postgres:doninha@localhost/petstore?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic("could not open connexion with DB")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/petstore/adapters/db/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	petRepository := adapters.NewPostgreSQLRepository(db)
	return app.Application{
		Commands: app.Commands{
			CreatePet: commands.NewCreatePetHandler(petRepository),
		},
		Queries: app.Queries{
			GetPet:   queries.NewGetPetHandler(petRepository),
			ListPets: queries.NewListPetsHandler(petRepository),
		},
	}
}

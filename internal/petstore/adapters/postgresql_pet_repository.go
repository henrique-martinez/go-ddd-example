package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"petstoreproject/internal/petstore/domain"

	"github.com/google/uuid"
)

type PostgreSQLPetRepository struct {
	db *sql.DB
}

var _ domain.PetRepository = (*PostgreSQLPetRepository)(nil)

func NewPostgreSQLRepository(db *sql.DB) PostgreSQLPetRepository {
	if db == nil {
		panic("nil db")
	}
	return PostgreSQLPetRepository{
		db: db,
	}
}

func (r PostgreSQLPetRepository) ListPets(ctx context.Context) ([]domain.Pet, error) {
	rows, err := r.db.QueryContext(ctx, "select id, name, tag from pet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []domain.Pet

	for rows.Next() {
		var id uuid.UUID
		var name string
		var tag sql.NullString
		if err := rows.Scan(&id, &name, &tag); err != nil {
			return pets, err
		}

		pets = append(pets, domain.Pet{
			Id:   id.String(),
			Name: name,
			Tag:  tag.String,
		})
	}
	if err = rows.Err(); err != nil {
		return pets, err
	}
	return pets, nil
}

func (r PostgreSQLPetRepository) CreatePet(ctx context.Context, pet domain.Pet) error {
	_, err := r.db.ExecContext(ctx, "insert into pet (id, name, tag) values ($1, $2, $3)", pet.Id, pet.Name, pet.Tag)
	return err
}

func (r PostgreSQLPetRepository) GetPet(ctx context.Context, id string) (domain.Pet, error) {
	var name string
	var tag sql.NullString
	if err := r.db.QueryRow("select name, tag from pet where id = $1", id).Scan(&name, &tag); err != nil {
		if err == sql.ErrNoRows {
			return domain.Pet{}, fmt.Errorf("pet %s not found", id)
		}
		return domain.Pet{}, fmt.Errorf("failed to fetch data: %v", err)
	}
	return domain.Pet{
		Id:   id,
		Name: name,
		Tag:  tag.String,
	}, nil
}

package ports

import (
	"net/http"
	"petstoreproject/internal/common/logs"
	"petstoreproject/internal/petstore/app"
	"petstoreproject/internal/petstore/domain"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func httpRespondWithError(err error, w http.ResponseWriter, r *http.Request, status int) {
	logs.GetLogEntry(r).Warn(err)
	w.WriteHeader(status)
	render.Respond(w, r, Error{
		Code:    status,
		Message: err.Error(),
	})
}

func petDomainToOpenAPI(pet domain.Pet) Pet {
	var tag *string
	if pet.Tag != "" {
		tag = &pet.Tag
	}
	return Pet{
		Id:   pet.Id,
		Name: pet.Name,
		Tag:  tag,
	}
}

// List all pets
// (GET /pets)
func (h HttpServer) ListPets(w http.ResponseWriter, r *http.Request, params ListPetsParams) {
	pets, err := h.app.Queries.ListPets.Handle(r.Context())
	if err != nil {
		httpRespondWithError(err, w, r, http.StatusInternalServerError)
		return
	}

	petsResponse := make([]Pet, len(pets))
	for idx, pet := range pets {
		petsResponse[idx] = petDomainToOpenAPI(pet)
	}

	render.Respond(w, r, petsResponse)
}

// Create a pet
// (POST /pets)
func (h HttpServer) CreatePets(w http.ResponseWriter, r *http.Request) {
	var body CreatePetsJSONBody

	if err := render.Decode(r, &body); err != nil {
		httpRespondWithError(err, w, r, http.StatusBadRequest)
		return
	}

	var tag string
	if body.Tag != nil {
		tag = *body.Tag
	}

	err := h.app.Commands.CreatePet.Handle(r.Context(), domain.Pet{
		Id:   uuid.NewString(),
		Name: body.Name,
		Tag:  tag,
	})
	if err != nil {
		httpRespondWithError(err, w, r, http.StatusInternalServerError)
		return
	}
}

// Info for a specific pet
// (GET /pets/{petId})
func (h HttpServer) ShowPetById(w http.ResponseWriter, r *http.Request, petId string) {
	pet, err := h.app.Queries.GetPet.Handle(r.Context(), petId)
	if err != nil {
		httpRespondWithError(err, w, r, http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, petDomainToOpenAPI(pet))
}

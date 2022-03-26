package ports

import (
	"net/http"
	"petstore/internal/petstore/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// List all pets
// (GET /pets)
func (h HttpServer) ListPets(w http.ResponseWriter, r *http.Request, params ListPetsParams) {
	panic("not implemented") // TODO: Implement
}

// Create a pet
// (POST /pets)
func (h HttpServer) CreatePets(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

// Info for a specific pet
// (GET /pets/{petId})
func (h HttpServer) ShowPetById(w http.ResponseWriter, r *http.Request, petId string) {
	panic("not implemented") // TODO: Implement
}

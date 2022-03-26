package main

import (
	"context"
	"net/http"
	"petstore/internal/common/server"
	"petstore/internal/petstore/ports"
	"petstore/internal/petstore/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(application),
			router,
		)
	})
}

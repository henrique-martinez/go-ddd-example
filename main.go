package main

import (
	"context"
	"net/http"
	"petstoreproject/internal/common/logs"
	"petstoreproject/internal/common/server"
	"petstoreproject/internal/petstore/ports"
	"petstoreproject/internal/petstore/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	logs.Init()
	application := service.NewApplication(ctx)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(application),
			router,
		)
	})
}

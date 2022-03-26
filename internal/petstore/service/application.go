package service

import (
	"context"
	"petstore/internal/petstore/app"
)

func NewApplication(ctx context.Context) app.Application {

	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}

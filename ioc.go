package main

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/config/database"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
	"github.com/sarulabs/di"
)

// NewIOC func
func NewIOC() di.Container {
	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: "database",
		Build: func(ctn di.Container) (interface{}, error) {
			db, err := database.ConfigureDatabase()
			return db, err
		},
	})

	builder.Add(di.Def{
		Name: "jwt",
		Build: func(ctn di.Container) (interface{}, error) {
			return config.ConfigureJWT(), nil
		},
	})

	builder.Add(di.Def{
		Name: "repository",
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewRepository(builder.Build()), nil
		},
	})

	builder.Add(di.Def{
		Name: "service",
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewService(builder.Build()), nil
		},
	})

	return builder.Build()
}

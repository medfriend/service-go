// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"service-go/controller"
	"service-go/repository"
	"service-go/service"
)

// Injectors from servicio.module.go:

func InitializeServicioModule(db *gorm.DB) *controller.ServiceController {
	servicioRepository := repository.NewServicioRepository(db)
	servicioService := service.NewServicioService(servicioRepository)
	serviceController := controller.NewServicioController(servicioService)
	return serviceController
}

// servicio.module.go:

var servicioSet = wire.NewSet(repository.NewServicioRepository, service.NewServicioService, controller.NewServicioController)

package module

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"service-go/controller"
	"service-go/repository"
	"service-go/service"
)

var servicioSet = wire.NewSet(
	repository.NewServicioRepository,
	service.NewServicioService,
	controller.NewServicioController)

func InitializeServicioModule(db *gorm.DB) *controller.ServiceController {
	wire.Build(servicioSet)
	return nil
}

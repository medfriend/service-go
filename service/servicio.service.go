package service

import (
	"service-go/entity"
	"service-go/repository"
)

type ServicioService interface {
	FindByPrefijo(prefijo string) (*[]entity.Service, error)
}

type ServicioServiceImpl struct {
	ServicioRepo repository.ServicioRepository
}

func NewServicioService(servicioRepo repository.ServicioRepository) ServicioService {
	return &ServicioServiceImpl{
		ServicioRepo: servicioRepo,
	}
}

func (s ServicioServiceImpl) FindByPrefijo(prefijo string) (*[]entity.Service, error) {
	return s.ServicioRepo.FindByPrefijo(prefijo)
}

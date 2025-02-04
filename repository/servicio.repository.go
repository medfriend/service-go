package repository

import (
	"github.com/medfriend/shared-commons-go/util/repository"
	"gorm.io/gorm"
	"service-go/entity"
)

type ServicioRepository interface {
	FindByPrefijo(prefijo string) (*[]entity.Service, error)
}

type ServicioRepositoryImpl struct {
	Base repository.BaseRepository[entity.Service]
}

func NewServicioRepository(db *gorm.DB) ServicioRepository {
	return &ServicioRepositoryImpl{
		Base: repository.BaseRepository[entity.Service]{DB: db},
	}
}

func (s ServicioRepositoryImpl) FindByPrefijo(prefijo string) (*[]entity.Service, error) {
	return s.Base.FindAnyField(
		[]string{"prefijo"},
		prefijo)
}

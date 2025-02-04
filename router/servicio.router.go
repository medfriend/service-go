package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"service-go/module"
)

func NewServicioRouter(router *gin.RouterGroup, db *gorm.DB) {
	servicioController := module.InitializeServicioModule(db)
	routerGroup := router.Group("auth")
	routerGroup.POST("/{prefijo}", servicioController.GetServiceByPrefijo)
}

func init() {
	RegisterRouter(NewServicioRouter)
}

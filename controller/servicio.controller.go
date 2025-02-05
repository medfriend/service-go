package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/medfriend/shared-commons-go/util/controller"
	"service-go/service"
)

type ServiceController struct {
	ServicioService service.ServicioService
}

func NewServicioController(servicioService service.ServicioService) *ServiceController {
	return &ServiceController{
		ServicioService: servicioService,
	}
}

// GetServiceByPrefijo obtener servicio por medio de prefijo
// @Summary obtener prefijo
// @Description obtener servicio.
// @Tags servicio
// @Accept json
// @Produce json
// @Param   prefijo  path  string true  "prefijo o abreviatura"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /servicio/{prefijo} [get]
func (ctrl *ServiceController) GetServiceByPrefijo(c *gin.Context) {
	servicio, err := ctrl.ServicioService.FindByPrefijo(c.Param("prefijo"))
	if err != nil {
	}
	controller.HandlerFoundSuccess(c, err, "servicio")
	controller.HandlerFound(c, servicio)
}

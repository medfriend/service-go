package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
)

func InitHttpServer(taskQueue chan *http.Request, db *gorm.DB, serviceInfo map[string]string) {
	r := gin.Default()

	api := r.Group(serviceInfo["SERVICE_NAME"])

	fmt.Println(taskQueue)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(fmt.Sprintf(":%s", serviceInfo["SERVICE_PORT"]))

	fmt.Println(api)

	if err != nil {
		return
	}
}

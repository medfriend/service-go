package httpServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"net/http"
	_ "service-go/docs"
	"service-go/router"
)

func InitHttpServer(taskQueue chan *http.Request, db *gorm.DB, serviceInfo map[string]string) {
	r := gin.Default()

	api := r.Group(serviceInfo["SERVICE_NAME"])

	fmt.Println(taskQueue)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.InitializeAllRouters(api, db)

	err := r.Run(fmt.Sprintf(":%s", serviceInfo["SERVICE_PORT"]))

	if err != nil {
		return
	}
}

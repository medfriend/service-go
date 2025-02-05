package main

// @title           medfri-service
// @version         1.0
// @description     micro de servicios.

// @host            localhost:9020
// @BasePath        /medfri-service

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ingresa "Bearer {token}" para autenticar.

// @contact.name    Soporte de API
// @contact.url     http://www.soporte-api.com
// @contact.email   soporte@api.com

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

import (
	"encoding/json"
	"fmt"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	gormUtil "github.com/medfriend/shared-commons-go/util/gorm"
	"github.com/medfriend/shared-commons-go/util/worker"
	"gorm.io/gorm"
	"net/http"
	"os"
	"runtime"
	"service-go/httpServer"
)

var db *gorm.DB

func main() {
	env.LoadEnv()

	consulClient := consul.ConnectToConsulKey("", "SERVICE")
	serviceInfo, _ := consul.GetKeyValue(consulClient, "SERVICE")

	var result map[string]string
	err := json.Unmarshal([]byte(serviceInfo), &result)

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	initDB, err := gormUtil.InitDB(
		db,
		consulClient,
		os.Getenv("SERVICE_STATUS"),
		"SERVICE",
	)

	httpServer.InitHttpServer(taskQueue, initDB, result)

	if err != nil {
		return
	}
}

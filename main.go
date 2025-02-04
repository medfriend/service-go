package main

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

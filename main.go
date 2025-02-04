package main

import (
	"fmt"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	gormUtil "github.com/medfriend/shared-commons-go/util/gorm"
	"github.com/medfriend/shared-commons-go/util/worker"
	"gorm.io/gorm"
	"net/http"
	"os"
	"runtime"
)

var db *gorm.DB

func main() {
	env.LoadEnv()
	consulClient := consul.ConnectToConsulKey("", "SERVICE")

	fmt.Println(consulClient)

	numCPUs := runtime.NumCPU()

	fmt.Printf("Detected %d CPUs, creating %d workers\n", numCPUs, numCPUs)

	taskQueue := make(chan *http.Request, 100)

	stop := make(chan struct{})

	worker.CreateWorkers(numCPUs, stop, taskQueue)

	_, err := gormUtil.InitDB(
		db,
		consulClient,
		os.Getenv("SERVICE_STATUS"),
		"SERVICE",
	)

	if err != nil {
		return
	}
}

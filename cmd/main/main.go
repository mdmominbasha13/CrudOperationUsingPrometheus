package main

import (
	"github.com/CrudOperationUsingPrometheus/pkg/logger"
	"github.com/CrudOperationUsingPrometheus/pkg/router"

	"github.com/CrudOperationUsingPrometheus/pkg/config"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()

	logger.InfoLn("Logger Initialized successfully")

	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}

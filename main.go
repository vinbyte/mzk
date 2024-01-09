package main

import (
	"github.com/vinbyte/mzk/configs"
	"github.com/vinbyte/mzk/shared/logger"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire
var config *configs.Config

func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize config
	config = configs.Get()

	// Set desired log level
	logger.SetLogLevel(config)

	// Wire everything up
	http := InitializeService()

	// Run server
	http.SetupAndServe()
}

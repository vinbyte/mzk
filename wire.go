//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/vinbyte/mzk/configs"
	"github.com/vinbyte/mzk/infras"
	"github.com/vinbyte/mzk/transport/http"
	"github.com/vinbyte/mzk/transport/http/router"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvidePostgresConn,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	// wire.Struct(new(router.DomainHandlers), "FooBarBazHandler"),
	// handlers.ProvideFooBarBazHandler,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// middleware
		// authMiddleware,
		// domains
		// domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}

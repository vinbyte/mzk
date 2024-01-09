//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/vinbyte/mzk/configs"
	"github.com/vinbyte/mzk/infras"
	"github.com/vinbyte/mzk/internal/domain/student"
	"github.com/vinbyte/mzk/internal/handlers"
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
	wire.Struct(new(router.DomainHandlers), "StudentHandler"),
	handlers.ProvideStudentHandler,
	router.ProvideRouter,
)

// Wiring for domain Student.
var domainStudent = wire.NewSet(
	// StudentService interface and implementation
	student.ProvideStudentServiceImpl,
	wire.Bind(new(student.StudentService), new(*student.StudentServiceImpl)),
	// StudentRepository interface and implementation
	student.ProvideStudentRepositoryPostgres,
	wire.Bind(new(student.StudentRepository), new(*student.StudentRepositoryPostgres)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainStudent,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}

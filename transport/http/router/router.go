package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/vinbyte/mzk/internal/handlers"
)

// DomainHandlers is a struct that contains all domain-specific handlers.
type DomainHandlers struct {
	StudentHandler handlers.StudentHandler
}

// Router is the router struct containing handlers.
type Router struct {
	DomainHandlers DomainHandlers
}

// ProvideRouter is the provider function for this router.
func ProvideRouter(domainHandlers DomainHandlers) Router {
	return Router{
		DomainHandlers: domainHandlers,
	}
}

// SetupRoutes sets up all routing for this app.
func (r *Router) SetupRoutes(mux *chi.Mux) {
	mux.Route("/v1", func(rc chi.Router) {
		r.DomainHandlers.StudentHandler.Router(rc)
	})
}

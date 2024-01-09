package router

// Router is the router struct containing handlers.
type Router struct{}

// ProvideRouter is the provider function for this router.
func ProvideRouter() Router {
	return Router{}
}

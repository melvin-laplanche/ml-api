package router

type RouteHandler func(*Request)

// Endpoint represents an HTTP endpoint
type Endpoint struct {
	Verb    string
	Path    string
	Auth    RouteAuth
	Handler RouteHandler
	Params  interface{}
}

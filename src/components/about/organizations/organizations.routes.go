package organizations

import (
	"github.com/Nivl/go-rest-tools/router"
	"github.com/gorilla/mux"
)

// Contains the index of all Endpoints
const (
	EndpointAdd = iota
	EndpointDelete
	EndpointUploadLogo
)

// Endpoints is a list of endpoints for this components
var Endpoints = router.Endpoints{
	EndpointAdd:        addEndpoint,
	EndpointDelete:     deleteEndpoint,
	EndpointUploadLogo: uploadLogoEndpoint,
}

// SetRoutes is used to set all the routes of the article
func SetRoutes(r *mux.Router) {
	Endpoints.Activate(r)
}
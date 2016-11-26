package articles

import (
	"github.com/gorilla/mux"
	"github.com/melvin-laplanche/ml-api/src/router"
)

// Indexes of all different endpoints
const (
	EndpointAdd = iota
	EndpointList
	EndpointUserList
)

// Endpoints contains the list of endpoints for this component
var Endpoints = router.Endpoints{
	EndpointAdd: {
		Verb:    "POST",
		Path:    "/articles",
		Handler: HandlerAdd,
		Auth:    router.LoggedUser,
		Params:  &HandlerAddParams{},
	},
	EndpointList: {
		Verb:    "GET",
		Path:    "/articles",
		Handler: HandlerList,
		Auth:    nil,
	},
	EndpointUserList: {
		Verb:    "GET",
		Prefix:  "/users/{user_id}",
		Path:    "/articles",
		Handler: HandlerListForUser,
		Auth:    nil,
		Params:  &HandlerListForUserParams{},
	},
}

// SetRoutes is used to set all the routes of the article
func SetRoutes(baseURI string, r *mux.Router) {
	Endpoints.Activate(baseURI, r)
}

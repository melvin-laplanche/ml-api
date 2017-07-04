package sessions

import (
	"github.com/Nivl/go-rest-tools/network/http/httperr"
	"github.com/Nivl/go-rest-tools/router"
	"github.com/Nivl/go-rest-tools/router/guard"
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/Nivl/go-rest-tools/storage/db"
)

var addEndpoint = &router.Endpoint{
	Verb:    "POST",
	Path:    "/sessions",
	Handler: Add,
	Guard: &guard.Guard{
		ParamStruct: &AddParams{},
	},
}

// AddParams represent the request params accepted by HandlerAdd
type AddParams struct {
	Email    string `from:"form" json:"email" params:"required,trim"`
	Password string `from:"form" json:"password" params:"required,trim"`
}

// Add represents an API handler to create a new user session
func Add(req router.HTTPRequest, deps *router.Dependencies) error {
	params := req.Params().(*AddParams)

	var user auth.User
	stmt := "SELECT * FROM users WHERE email=$1 LIMIT 1"
	err := db.Get(deps.DB, &user, stmt, params.Email)
	if err != nil {
		return err
	}

	if user.ID == "" || !auth.IsPasswordValid(user.Password, params.Password) {
		return httperr.NewBadRequest("Bad email/password")
	}

	s := &auth.Session{
		UserID: user.ID,
	}
	if err := s.Save(deps.DB); err != nil {
		return err
	}

	req.Response().Created(NewPayload(s))
	return nil
}
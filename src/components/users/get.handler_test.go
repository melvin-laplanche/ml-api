package users_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/Nivl/go-rest-tools/router"
	"github.com/Nivl/go-rest-tools/router/guard/testguard"
	"github.com/Nivl/go-rest-tools/router/mockrouter"
	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/Nivl/go-rest-tools/storage/db/mockdb"
	"github.com/Nivl/go-rest-tools/types/apierror"
	"github.com/melvin-laplanche/ml-api/src/components/users"
	"github.com/melvin-laplanche/ml-api/src/components/users/testusers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetInvalidParams(t *testing.T) {
	testCases := []testguard.InvalidParamsTestCase{
		{
			Description: "Should fail on missing ID",
			MsgMatch:    params.ErrMsgMissingParameter,
			FieldName:   "id",
			Sources: map[string]url.Values{
				"url": url.Values{},
			},
		},
		{
			Description: "Should fail on invalid ID",
			MsgMatch:    params.ErrMsgInvalidUUID,
			FieldName:   "id",
			Sources: map[string]url.Values{
				"url": url.Values{
					"id": []string{"not-a-uuid"},
				},
			},
		},
	}

	g := users.Endpoints[users.EndpointGet].Guard
	testguard.InvalidParams(t, g, testCases)
}

func TestGetValidParams(t *testing.T) {
	testCases := []struct {
		description string
		sources     map[string]url.Values
	}{
		{
			"Should work with a valid ID",
			map[string]url.Values{
				"url": url.Values{
					"id": []string{"48d0c8b8-d7a3-4855-9d90-29a06ef474b0"},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			endpts := users.Endpoints[users.EndpointGet]
			data, err := endpts.Guard.ParseParams(tc.sources, nil)
			assert.NoError(t, err)

			if data != nil {
				p := data.(*users.GetParams)
				assert.Equal(t, tc.sources["url"].Get("id"), p.ID)
			}
		})
	}
}

func TestGetOthersData(t *testing.T) {
	userToGet := testusers.NewProfile()

	handlerParams := &users.GetParams{
		ID: userToGet.ID,
	}
	requester := &auth.User{
		ID: "48d0c8b8-d7a3-4855-9d90-29a06ef474b0",
	}

	// Mock the database & add expectations
	mockDB := &mockdb.Connection{}
	mockDB.ExpectGet("*users.Profile", func(args mock.Arguments) {
		user := args.Get(0).(*users.Profile)
		*user = *userToGet
	})

	// Mock the response & add expectations
	res := new(mockrouter.HTTPResponse)
	res.ExpectOk("*users.ProfilePayload", func(args mock.Arguments) {
		pld := args.Get(0).(*users.ProfilePayload)
		assert.Equal(t, userToGet.User.ID, pld.ID, "The user ID should not have changed")
		assert.Equal(t, userToGet.Name, pld.Name, "Name should not have changed")
		assert.Equal(t, *userToGet.LinkedIn, pld.LinkedIn, "the LinkedIn id should not have changed")
		assert.Empty(t, pld.Email, "the email should not be returned to anyone")
		assert.False(t, pld.IsAdmin, "user should not be an admin")
	})

	// Mock the request & add expectations
	req := new(mockrouter.HTTPRequest)
	req.On("Response").Return(res)
	req.On("Params").Return(handlerParams)
	req.On("User").Return(requester)

	// call the handler
	err := users.Get(req, &router.Dependencies{DB: mockDB})

	// Assert everything
	assert.NoError(t, err, "the handler should not have fail")
	mockDB.AssertExpectations(t)
	req.AssertExpectations(t)
	res.AssertExpectations(t)
}

func TestGetOwnData(t *testing.T) {
	handlerParams := &users.GetParams{
		ID: "0c2f0713-3f9b-4657-9cdd-2b4ed1f214e9",
	}
	requester := &auth.User{
		ID:      handlerParams.ID,
		Name:    "user name",
		Email:   "email@domain.tld",
		IsAdmin: false,
	}

	// Mock the database & add expectations
	mockDB := &mockdb.Connection{}
	mockDB.ExpectGet("*users.Profile", func(args mock.Arguments) {
		profile := args.Get(0).(*users.Profile)
		*profile = *(testusers.NewProfile())
		profile.User = requester
		profile.UserID = requester.ID
	})

	// Mock the response & add expectations
	res := new(mockrouter.HTTPResponse)
	res.ExpectOk("*users.ProfilePayload", func(args mock.Arguments) {
		pld := args.Get(0).(*users.ProfilePayload)
		assert.Equal(t, requester.ID, pld.ID, "ID should have not changed")
		assert.Equal(t, requester.Name, pld.Name, "Name should have not changed")
		assert.Equal(t, requester.Email, pld.Email, "the email should be returned")
		assert.False(t, pld.IsAdmin, "user should not be an admin")
	})

	// Mock the request & add expectations
	req := new(mockrouter.HTTPRequest)
	req.On("Response").Return(res)
	req.On("Params").Return(handlerParams)
	req.On("User").Return(requester)

	// call the handler
	err := users.Get(req, &router.Dependencies{DB: mockDB})

	// Assert everything
	assert.NoError(t, err, "the handler should not have fail")
	req.AssertExpectations(t)
	res.AssertExpectations(t)
}

func TestGetUnexistingUser(t *testing.T) {
	handlerParams := &users.GetParams{
		ID: "0c2f0713-3f9b-4657-9cdd-2b4ed1f214e9",
	}

	// Mock the database & add expectations
	mockDB := &mockdb.Connection{}
	mockDB.ExpectGetNotFound("*users.Profile")

	// Mock the request & add expectations
	req := new(mockrouter.HTTPRequest)
	req.On("Params").Return(handlerParams)

	// call the handler
	err := users.Get(req, &router.Dependencies{DB: mockDB})

	// Assert everything
	assert.Error(t, err, "the handler should have fail")
	mockDB.AssertExpectations(t)
	req.AssertExpectations(t)

	httpErr := apierror.Convert(err)
	assert.Equal(t, http.StatusNotFound, httpErr.HTTPStatus())
}

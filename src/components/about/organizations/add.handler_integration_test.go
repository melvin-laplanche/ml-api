// +build integration

package organizations_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nivl/go-rest-tools/dependencies"
	"github.com/Nivl/go-rest-tools/network/http/httptests"
	"github.com/Nivl/go-rest-tools/primitives/models/lifecycle"
	"github.com/Nivl/go-rest-tools/primitives/ptrs"
	"github.com/Nivl/go-rest-tools/router"
	"github.com/Nivl/go-rest-tools/security/auth/testauth"
	"github.com/dchest/uniuri"
	"github.com/melvin-laplanche/ml-api/src/components/about/organizations"
	"github.com/melvin-laplanche/ml-api/src/components/about/organizations/testorganizations"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	dbCon := dependencies.DB

	defer lifecycle.PurgeModels(t, dbCon)
	_, admSession := testauth.NewAdminAuth(t, dbCon)
	adminAuth := httptests.NewRequestAuth(admSession)

	tests := []struct {
		description string
		code        int
		params      *organizations.AddParams
	}{
		{
			"Valid Request should work",
			http.StatusCreated,
			&organizations.AddParams{Name: uniuri.New(), Website: ptrs.NewString("http://google.com")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			rec := callAdd(t, tc.params, adminAuth)
			assert.Equal(t, tc.code, rec.Code)

			if rec.Code == http.StatusCreated {
				org := &organizations.Organization{}
				if err := json.NewDecoder(rec.Body).Decode(org); err != nil {
					t.Fatal(err)
				}

				assert.NotEmpty(t, org.ID)
				assert.Equal(t, tc.params.Name, org.Name)
				assert.Nil(t, org.ShortName)
				assert.NotNil(t, org.Website)
				assert.Equal(t, *tc.params.Website, *org.Website)

				// clean the test
				org.Delete(dbCon)
			}
		})
	}
}

func TestIntegrationAddConflictName(t *testing.T) {
	dbCon := dependencies.DB
	defer lifecycle.PurgeModels(t, dbCon)

	_, admSession := testauth.NewAdminAuth(t, dbCon)
	adminAuth := httptests.NewRequestAuth(admSession)
	org := testorganizations.NewOrganization(t, dbCon, nil)

	p := &organizations.AddParams{
		Name: org.Name,
	}

	rec := callAdd(t, p, adminAuth)
	assert.Equal(t, http.StatusConflict, rec.Code)

	pld := &router.ResponseError{}
	if err := json.NewDecoder(rec.Body).Decode(pld); err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, org.ID)
	assert.Equal(t, "already exists", pld.Error)
	assert.Equal(t, "name", pld.Field)
}

func TestIntegrationAddConflictShortName(t *testing.T) {
	dbCon := dependencies.DB
	defer lifecycle.PurgeModels(t, dbCon)

	_, admSession := testauth.NewAdminAuth(t, dbCon)
	adminAuth := httptests.NewRequestAuth(admSession)
	org := testorganizations.NewOrganization(t, dbCon, nil)

	p := &organizations.AddParams{
		Name:      uniuri.New(),
		ShortName: org.ShortName,
	}

	rec := callAdd(t, p, adminAuth)
	assert.Equal(t, http.StatusConflict, rec.Code)

	pld := &router.ResponseError{}
	if err := json.NewDecoder(rec.Body).Decode(pld); err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, org.ID)
	assert.Equal(t, "already exists", pld.Error)
	assert.Equal(t, "short_name", pld.Field)
}

func callAdd(t *testing.T, params *organizations.AddParams, auth *httptests.RequestAuth) *httptest.ResponseRecorder {
	ri := &httptests.RequestInfo{
		Endpoint: organizations.Endpoints[organizations.EndpointAdd],
		Params:   params,
		Auth:     auth,
	}

	return httptests.NewRequest(t, ri)
}
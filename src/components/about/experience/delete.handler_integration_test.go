// +build integration

package experience_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nivl/go-rest-tools/dependencies"
	"github.com/Nivl/go-rest-tools/network/http/httptests"
	"github.com/Nivl/go-rest-tools/security/auth/testauth"
	"github.com/Nivl/go-rest-tools/storage/db"
	"github.com/Nivl/go-rest-tools/types/models/lifecycle"
	"github.com/melvin-laplanche/ml-api/src/components/about/experience"
	"github.com/melvin-laplanche/ml-api/src/components/about/experience/testexperience"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationDeleteHappyPath(t *testing.T) {
	dbCon := dependencies.DB

	defer lifecycle.PurgeModels(t, dbCon)
	_, admSession := testauth.NewAdminAuth(t, dbCon)
	adminAuth := httptests.NewRequestAuth(admSession)
	basicExp := testexperience.NewPersisted(t, dbCon, nil)
	trashedExp := testexperience.NewPersisted(t, dbCon, &experience.Experience{
		DeletedAt: db.Now(),
	})

	tests := []struct {
		description string
		code        int
		params      *experience.DeleteParams
	}{
		{
			"Valid request should work",
			http.StatusNoContent,
			&experience.DeleteParams{ID: basicExp.ID},
		},
		{
			"trashed exp should work",
			http.StatusNoContent,
			&experience.DeleteParams{ID: trashedExp.ID},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			rec := callDelete(t, tc.params, adminAuth)
			assert.Equal(t, tc.code, rec.Code)

			if rec.Code == http.StatusNoContent {
				exists, err := experience.Exists(dbCon, tc.params.ID)
				assert.NoError(t, err, "Exists() should have not failed")
				assert.False(t, exists, "the organization should no longer exists")
			}
		})
	}
}

func callDelete(t *testing.T, params *experience.DeleteParams, auth *httptests.RequestAuth) *httptest.ResponseRecorder {
	ri := &httptests.RequestInfo{
		Endpoint: experience.Endpoints[experience.EndpointDelete],
		Params:   params,
		Auth:     auth,
	}
	return httptests.NewRequest(t, ri)
}
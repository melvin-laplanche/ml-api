package organizations

import (
	"fmt"
	"net/http"

	"github.com/Nivl/go-params/formfile"
	"github.com/Nivl/go-rest-tools/router"
	"github.com/Nivl/go-rest-tools/router/guard"
	"github.com/Nivl/go-rest-tools/storage/filestorage"
	"github.com/Nivl/go-types/filetype"
	"github.com/Nivl/go-types/ptrs"
)

var uploadLogoEndpoint = &router.Endpoint{
	Verb:    http.MethodPut,
	Path:    "/about/organizations/{id}/logo",
	Handler: UploadLogo,
	Guard: &guard.Guard{
		Auth:        guard.AdminAccess,
		ParamStruct: &UploadLogoParams{},
	},
}

// UploadLogoParams represents the params accepted by the UploadLogo endpoint
type UploadLogoParams struct {
	ID   string             `from:"url" json:"id" params:"required,uuid"`
	Logo *formfile.FormFile `from:"file" json:"logo" params:"required,image"`
}

// UploadLogo is an endpoint used to upload a logo
func UploadLogo(req router.HTTPRequest, deps *router.Dependencies) error {
	params := req.Params().(*UploadLogoParams)

	org, err := GetByID(deps.DB, params.ID)
	if err != nil {
		return err
	}

	// we use the shasum as filename that way we don't re-upload the same
	// image twice
	shasum, err := filetype.SHA256Sum(params.Logo.File)
	if err != nil {
		return err
	}

	fileDest := fmt.Sprintf("about/organizations/%s", shasum)
	_, url, err := deps.Storage.WriteIfNotExist(params.Logo.File, fileDest)
	if err != nil {
		return err
	}

	// not a big deal if that fails
	deps.Storage.SetAttributes(fileDest, &filestorage.UpdatableFileAttributes{
		ContentType: params.Logo.Mime,
	})

	// Save the new logo
	org.Logo = ptrs.NewString(url)
	if err = org.Update(deps.DB); err != nil {
		return err
	}

	return req.Response().Ok(org.ExportPrivate())
}

package articles

import (
	"github.com/gosimple/slug"
	"github.com/melvin-laplanche/ml-api/src/db"
	"github.com/melvin-laplanche/ml-api/src/ptrs"
	"github.com/melvin-laplanche/ml-api/src/router"
)

// HandlerAddParams lists the params allowed by HandlerAdd
type HandlerAddParams struct {
	Title       string `from:"form" json:"title,omitempty" params:"required,trim"`
	Subtitle    string `from:"form" json:"subtitle,omitempty"`
	Description string `from:"form" json:"description,omitempty"`
	Content     string `from:"form" json:"content,omitempty"`
}

// HandlerAdd represents an API handler to add a new article
func HandlerAdd(req *router.Request) {
	params := req.Params.(*HandlerAddParams)

	content := &Content{
		Title:       params.Title,
		Subtitle:    params.Subtitle,
		Content:     params.Content,
		Description: params.Description,
		IsCurrent:   ptrs.NewBool(true),
	}

	a := &Article{
		Slug:   slug.Make(content.Title),
		UserID: req.User.ID,
		User:   req.User,
	}

	tx, err := db.Con().Beginx()
	if err != nil {
		req.Error(err)
		return
	}
	defer tx.Rollback()

	if err := a.SaveTx(tx); err != nil {
		req.Error(err)
		return
	}

	content.ArticleID = a.ID
	if err := content.SaveTx(tx); err != nil {
		req.Error(err)
		return
	}

	if err := tx.Commit(); err != nil {
		req.Error(err)
		return
	}

	a.Content = content
	req.Created(a.PublicExport())
}

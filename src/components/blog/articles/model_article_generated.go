package articles

// Code generated by api-cli; DO NOT EDIT\n

import (
	"errors"
	"fmt"

	"github.com/melvin-laplanche/ml-api/src/apierror"
	"github.com/melvin-laplanche/ml-api/src/app"
	"github.com/melvin-laplanche/ml-api/src/db"
	uuid "github.com/satori/go.uuid"
)

// JoinSQL returns a string ready to be embed in a JOIN query
func JoinSQL(prefix string) string {
	fields := []string{"id", "slug", "created_at", "updated_at", "deleted_at", "published_at", "user_id"}
	output := ""

	for i, field := range fields {
		if i != 0 {
			output += ", "
		}

		fullName := fmt.Sprintf("%s.%s", prefix, field)
		output += fmt.Sprintf("%s \"%s\"", fullName, fullName)
	}

	return output
}

// Save creates or updates the article depending on the value of the id
func (a *Article) Save() error {
	if a == nil {
		return apierror.NewServerError("article is not instanced")
	}

	if a.ID == "" {
		return a.Create()
	}

	return a.Update()
}

// doCreate persists an object in the database
func (a *Article) doCreate() error {
	if a == nil {
		return errors.New("article not instanced")
	}

	a.ID = uuid.NewV4().String()
	a.CreatedAt = db.Now()
	a.UpdatedAt = db.Now()

	stmt := "INSERT INTO blog_articles (id, slug, created_at, updated_at, deleted_at, published_at, user_id) VALUES (:id, :slug, :created_at, :updated_at, :deleted_at, :published_at, :user_id)"
	fmt.Printf("\n\n Query: %s\nData %#v \n\n", stmt, a)
	_, err := app.GetContext().SQL.NamedExec(stmt, a)
	return err
}

// doUpdate updates an object in the database
func (a *Article) doUpdate() error {
	if a == nil {
		return apierror.NewServerError("article is not instanced")
	}

	if a.ID == "" {
		return apierror.NewServerError("cannot update a non-persisted article")
	}

	a.UpdatedAt = db.Now()

	stmt := "UPDATE blog_articles SET id = $1, slug = $2, created_at = $3, updated_at = $4, deleted_at = $5, published_at = $6, user_id = $7 WHERE id=$8"
	_, err := app.GetContext().SQL.Exec(stmt, a.ID, a.Slug, a.CreatedAt, a.UpdatedAt, a.DeletedAt, a.PublishedAt, a.UserID, a.ID)
	return err
}

// FullyDelete removes an object from the database
func (a *Article) FullyDelete() error {
	if a == nil {
		return errors.New("article not instanced")
	}

	if a.ID == "" {
		return errors.New("article has not been saved")
	}

	_, err := sql().Exec("DELETE FROM blog_articles WHERE id=$1", a.ID)
	return err
}

// Delete soft delete an object.
func (a *Article) Delete() error {
	return a.doDelete()
}

// doDelete performs a soft delete operation on an object
func (a *Article) doDelete() error {
	if a == nil {
		return apierror.NewServerError("article is not instanced")
	}

	if a.ID == "" {
		return apierror.NewServerError("cannot delete a non-persisted article")
	}

	a.DeletedAt = db.Now()

	stmt := "UPDATE blog_articles SET deleted_at = $2 WHERE id=$1"
	_, err := sql().Exec(stmt, a.ID, *a.DeletedAt)
	return err
}

// IsZero checks if the object is either nil or don't have an ID
func (a *Article) IsZero() bool {
	return a == nil || a.ID == ""
}

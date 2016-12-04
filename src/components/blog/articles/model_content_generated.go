package articles

// Code generated by api-cli; DO NOT EDIT\n

import (
	"errors"


	"github.com/Nivl/sqalx"
	"github.com/melvin-laplanche/ml-api/src/apierror"
	"github.com/melvin-laplanche/ml-api/src/db"
	uuid "github.com/satori/go.uuid"
)



// Save creates or updates the content depending on the value of the id
func (c *Content) Save() error {
	return c.SaveTx(db.Con())
}

// SaveTx creates or updates the article depending on the value of the id using
// a transaction
func (c *Content) SaveTx(tx sqalx.Node) error {
	if c == nil {
		return apierror.NewServerError("content is not instanced")
	}

	if c.ID == "" {
		return c.CreateTx(tx)
	}

	return c.UpdateTx(tx)
}

// Create persists a user in the database
func (c *Content) Create() error {
	return c.CreateTx(db.Con())
}

// Create persists a user in the database
func (c *Content) CreateTx(tx sqalx.Node) error {
	if c == nil {
		return apierror.NewServerError("content is not instanced")
	}

	if c.ID != "" {
		return apierror.NewServerError("cannot persist a content that already has an ID")
	}

	return c.doCreate(tx)
}

// doCreate persists an object in the database using a Node
func (c *Content) doCreate(tx sqalx.Node) error {
	if c == nil {
		return errors.New("content not instanced")
	}

	c.ID = uuid.NewV4().String()
	c.CreatedAt = db.Now()
	c.UpdatedAt = db.Now()

	stmt := "INSERT INTO blog_article_contents (id, created_at, updated_at, deleted_at, article_id, is_current, is_draft, title, content, subtitle, description) VALUES (:id, :created_at, :updated_at, :deleted_at, :article_id, :is_current, :is_draft, :title, :content, :subtitle, :description)"
	_, err := tx.NamedExec(stmt, c)

  return err
}

// Update updates most of the fields of a persisted content.
// Excluded fields are id, created_at, deleted_at, etc.
func (c *Content) Update() error {
	return c.UpdateTx(db.Con())
}

// Update updates most of the fields of a persisted content using a transaction
// Excluded fields are id, created_at, deleted_at, etc.
func (c *Content) UpdateTx(tx sqalx.Node) error {
	if c == nil {
		return apierror.NewServerError("content is not instanced")
	}

	if c.ID == "" {
		return apierror.NewServerError("cannot update a non-persisted content")
	}

	return c.doUpdate(tx)
}

// doUpdate updates an object in the database using an optional transaction
func (c *Content) doUpdate(tx sqalx.Node) error {
	if c == nil {
		return apierror.NewServerError("content is not instanced")
	}

	if c.ID == "" {
		return apierror.NewServerError("cannot update a non-persisted content")
	}

	c.UpdatedAt = db.Now()

	stmt := "UPDATE blog_article_contents SET id=:id, created_at=:created_at, updated_at=:updated_at, deleted_at=:deleted_at, article_id=:article_id, is_current=:is_current, is_draft=:is_draft, title=:title, content=:content, subtitle=:subtitle, description=:description WHERE id=:id"
	_, err := tx.NamedExec(stmt, c)

	return err
}

// FullyDelete removes an object from the database
func (c *Content) FullyDelete() error {
	return c.FullyDeleteTx(db.Con())
}

// FullyDeleteTx removes an object from the database using a transaction
func (c *Content) FullyDeleteTx(tx sqalx.Node) error {
	if c == nil {
		return errors.New("content not instanced")
	}

	if c.ID == "" {
		return errors.New("content has not been saved")
	}

	stmt := "DELETE FROM blog_article_contents WHERE id=$1"
	_, err := tx.Exec(stmt, c.ID)

	return err
}

// Delete soft delete an object.
func (c *Content) Delete() error {
	return c.DeleteTx(db.Con())
}

// DeleteTx soft delete an object using a transaction
func (c *Content) DeleteTx(tx sqalx.Node) error {
	return c.doDelete(tx)
}

// doDelete performs a soft delete operation on an object using an optional transaction
func (c *Content) doDelete(tx sqalx.Node) error {
	if c == nil {
		return apierror.NewServerError("content is not instanced")
	}

	if c.ID == "" {
		return apierror.NewServerError("cannot delete a non-persisted content")
	}

	c.DeletedAt = db.Now()

	stmt := "UPDATE blog_article_contents SET deleted_at = $2 WHERE id=$1"
	_, err := tx.Exec(stmt, c.ID, c.DeletedAt)
	return err
}

// IsZero checks if the object is either nil or don't have an ID
func (c *Content) IsZero() bool {
	return c == nil || c.ID == ""
}
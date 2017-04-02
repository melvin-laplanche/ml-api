package articles

// Code generated by api-cli; DO NOT EDIT\n

// // JoinSQL returns a string ready to be embed in a JOIN query
// func JoinSQL(prefix string) string {
// 	fields := []string{ "id", "slug", "created_at", "updated_at", "deleted_at", "published_at", "user_id" }
// 	output := ""

// 	for i, field := range fields {
// 		if i != 0 {
// 			output += ", "
// 		}

// 		fullName := fmt.Sprintf("%s.%s", prefix, field)
// 		output += fmt.Sprintf("%s \"%s\"", fullName, fullName)
// 	}

// 	return output
// }

// // Get finds and returns an active article by ID
// func Get(id string) (*Article, error) {
// 	a := &Article{}
// 	stmt := "SELECT * from blog_articles WHERE id=$1 and deleted_at IS NULL LIMIT 1"
// 	err := db.Get(a, stmt, id)
// 	// We want to return nil if a article is not found
// 	if a.ID == "" {
// 		return nil, err
// 	}
// 	return a, err
// }

// // Exists checks if  article by ID
// func Exists(id string) (bool, error) {
// 	exists := false
// 	stmt := "SELECT exists(SELECT 1 FROM blog_articles WHERE id=$1 and deleted_at IS NULL)"
// 	err := db.Con().Get(&exists, stmt, id)
// 	return exists, err
// }

// // Save creates or updates the article depending on the value of the id
// func (a *Article) Save() error {
// 	return a.SaveTx(db.Con())
// }

// // SaveTx creates or updates the article depending on the value of the id using
// // a transaction
// func (a *Article) SaveTx(tx sqalx.Node) error {
// 	if a == nil {
// 		return apierror.NewServerError("article is not instanced")
// 	}

// 	if a.ID == "" {
// 		return a.CreateTx(tx)
// 	}

// 	return a.UpdateTx(tx)
// }

// // Create persists a article in the database
// func (a *Article) Create() error {
// 	return a.CreateTx(db.Con())
// }

// // doCreate persists a article in the database using a Node
// func (a *Article) doCreate(tx sqalx.Node) error {
// 	if a == nil {
// 		return errors.New("article not instanced")
// 	}

// 	a.ID = uuid.NewV4().String()
// 	a.CreatedAt = db.Now()
// 	a.UpdatedAt = db.Now()

// 	stmt := "INSERT INTO blog_articles (id, slug, created_at, updated_at, deleted_at, published_at, user_id) VALUES (:id, :slug, :created_at, :updated_at, :deleted_at, :published_at, :user_id)"
// 	_, err := tx.NamedExec(stmt, a)

//   return err
// }

// // Update updates most of the fields of a persisted article.
// // Excluded fields are id, created_at, deleted_at, etc.
// func (a *Article) Update() error {
// 	return a.UpdateTx(db.Con())
// }

// // doUpdate updates a article in the database using an optional transaction
// func (a *Article) doUpdate(tx sqalx.Node) error {
// 	if a == nil {
// 		return apierror.NewServerError("article is not instanced")
// 	}

// 	if a.ID == "" {
// 		return apierror.NewServerError("cannot update a non-persisted article")
// 	}

// 	a.UpdatedAt = db.Now()

// 	stmt := "UPDATE blog_articles SET id=:id, slug=:slug, created_at=:created_at, updated_at=:updated_at, deleted_at=:deleted_at, published_at=:published_at, user_id=:user_id WHERE id=:id"
// 	_, err := tx.NamedExec(stmt, a)

// 	return err
// }

// // FullyDelete removes a article from the database
// func (a *Article) FullyDelete() error {
// 	return a.FullyDeleteTx(db.Con())
// }

// // FullyDeleteTx removes a article from the database using a transaction
// func (a *Article) FullyDeleteTx(tx sqalx.Node) error {
// 	if a == nil {
// 		return errors.New("article not instanced")
// 	}

// 	if a.ID == "" {
// 		return errors.New("article has not been saved")
// 	}

// 	stmt := "DELETE FROM blog_articles WHERE id=$1"
// 	_, err := tx.Exec(stmt, a.ID)

// 	return err
// }

// // Delete soft delete a article.
// func (a *Article) Delete() error {
// 	return a.DeleteTx(db.Con())
// }

// // DeleteTx soft delete a article using a transaction
// func (a *Article) DeleteTx(tx sqalx.Node) error {
// 	return a.doDelete(tx)
// }

// // doDelete performs a soft delete operation on a article using an optional transaction
// func (a *Article) doDelete(tx sqalx.Node) error {
// 	if a == nil {
// 		return apierror.NewServerError("article is not instanced")
// 	}

// 	if a.ID == "" {
// 		return apierror.NewServerError("cannot delete a non-persisted article")
// 	}

// 	a.DeletedAt = db.Now()

// 	stmt := "UPDATE blog_articles SET deleted_at = $2 WHERE id=$1"
// 	_, err := tx.Exec(stmt, a.ID, a.DeletedAt)
// 	return err
// }

// // IsZero checks if the object is either nil or don't have an ID
// func (a *Article) IsZero() bool {
// 	return a == nil || a.ID == ""
// }

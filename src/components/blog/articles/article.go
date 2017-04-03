package articles

import "github.com/Nivl/go-rest-tools/storage/db"

// Article is a structure representing an article that can be saved in the database
//go:generate api-cli generate model Article -t blog_articles -e CreateTx,UpdateTx
type Article struct {
	ID          string   `db:"id"`
	Slug        string   `db:"slug"`
	CreatedAt   *db.Time `db:"created_at"`
	UpdatedAt   *db.Time `db:"updated_at"`
	DeletedAt   *db.Time `db:"deleted_at"`
	PublishedAt *db.Time `db:"published_at"`
	UserID      string   `db:"user_id"`

	*Version `db:"version"`
}

// Articles represents a list of Articles
type Articles []Article

// // FetchDraft fetches the article's draft from the database and attached
// // it to the current object
// func (a *Article) FetchDraft() error {
// 	if a == nil {
// 		return apierror.NewServerError("article not instanced")
// 	}

// 	d := &Draft{}
// 	stmt := `SELECT *
// 					FROM blog_article_contents
// 					WHERE deleted_at IS NULL
//             AND article_id = $1
// 						AND is_draft IS true`

// 	if err := db.Get(d, stmt, a.ID); err != nil {
// 		return err
// 	}

// 	if !d.IsZero() {
// 		a.Draft = d
// 	}

// 	return nil
// }

// // CreateTx persists an article in the database
// func (a *Article) CreateTx(tx sqalx.Node) error {
// 	if a == nil {
// 		return apierror.NewServerError("article not instanced")
// 	}

// 	if a.Slug == "" {
// 		return apierror.NewServerError("cannot persist an article with no slug")
// 	}

// 	// To prevent duplicates on the slug, we'll retry the insert() up to 10 times
// 	originalSlug := a.Slug
// 	var err error
// 	for i := 0; i < 10; i++ {
// 		// We create a savepoint so we can rollback in case the request fails
// 		// (a rollback to a savepoint is required inside a transaction)
// 		savePoint, err := tx.Beginx()
// 		if err != nil {
// 			fmt.Printf("\n\n Failed to create savepoint: %d \n\n", i)
// 			return err
// 		}

// 		// we do the insert and treat the error if theres one
// 		if err = a.doCreate(savePoint); err != nil {
// 			savePoint.Rollback()

// 			if db.SQLIsDup(err) == false {
// 				return apierror.NewServerError(err.Error())
// 			}

// 			// In case of duplicate we'll add "-X" at the end of the slug, where X is
// 			// a number
// 			a.Slug = fmt.Sprintf("%s-%d", originalSlug, i)
// 		} else {
// 			savePoint.Commit()
// 			// everything went well
// 			return nil
// 		}
// 	}

// 	// after 10 try we just return an error
// 	return apierror.NewConflict(err.Error())
// }

// // UpdateTx updates most of the fields of a persisted user.
// // Excluded fields are id, created_at, deleted_at
// func (a *Article) UpdateTx(tx sqalx.Node) error {
// 	return nil
// }
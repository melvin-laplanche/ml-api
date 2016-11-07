package articles

import (
	"github.com/Nivl/api.melvin.la/api/auth"
	"github.com/Nivl/api.melvin.la/api/router"
)

// HandlerList represents a API handler to get a list of articles
func HandlerList(req *router.Request) {
	arts := []Article{}

	stmt := `SELECT articles.*, ` + auth.UserForeignSelect("users") + `
					FROM blog_articles articles
					LEFT JOIN users ON users.id = articles.user_id
					WHERE articles.deleted_at IS NULL
					ORDER BY articles.created_at`
	if err := sql().Select(&arts, stmt); err != nil {
		req.Error(err)
		return
	}

	req.Ok(NewPayloadFromModels(&arts))
}

package articles

// // HandlerList represents a API handler to get a list of articles
// func HandlerList(req *router.Request) error {
// 	arts := Articles{}
// 	stmt := `SELECT articles.*,
//                   ` + auth.UserJoinSQL("users") + `,
//                   ` + ContentJoinSQL("content") + `
// 					FROM blog_articles articles
// 					JOIN users ON users.id = articles.user_id
// 					JOIN blog_article_contents content ON content.article_id = articles.id
// 					WHERE articles.deleted_at IS NULL
// 						AND articles.published_at IS NOT NULL
// 						AND content.is_current IS TRUE
// 					ORDER BY articles.created_at`
// 	if err := db.Con().Select(&arts, stmt); err != nil {
// 		return err
// 	}

// 	req.Ok(arts.PublicExport())
// 	return nil
// }

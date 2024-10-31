package repositories

import (
	"DevBookAPI/src/models"
	"database/sql"
	"time"
)

type Posts struct {
	db *sql.DB
}

func NewRepositoryPosts(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) FindOnePost(postId uint64) (models.Posts, error) {
	stmt, err := repository.db.Query("SELECT id, title, content, like_count, created_at FROM posts WHERE id = ?", postId)

	if err != nil {
		return models.Posts{}, err
	}
	defer stmt.Close()

	post := models.Posts{}
	if stmt.Next() {
		err := stmt.Scan(&post.Id, &post.Title, &post.Content, &post.LikeCount, &post.CreatedAt)
		if err != nil {
			return models.Posts{}, err
		}
	}

	return post, nil
}
func (repository Posts) Create(post models.Posts) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO posts (id, title, content, author_id, author_nick, like_count, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(post.Id, post.Title, post.Content, post.AuthorId, post.AuthorNick, post.LikeCount, time.Now())
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil

}
